package biz

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ifuryst/lol"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/ai/agent"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/helper/auth"
)

type Trip struct {
	tripRepo repository.TripRepo
	poiRepo  repository.PointsOfInterestRepo
	cityRepo repository.CityRepo
	logger   *zap.SugaredLogger
	agentHub *agent.Hub
}

func NewTrip(tripRepo repository.TripRepo, poiRepo repository.PointsOfInterestRepo, cityRepo repository.CityRepo, logger *zap.SugaredLogger, agentHub *agent.Hub) *Trip {
	return &Trip{
		tripRepo: tripRepo,
		poiRepo:  poiRepo,
		cityRepo: cityRepo,
		logger:   logger.Named("biz.tripRepo"),
		agentHub: agentHub,
	}
}

func (b *Trip) createTripFromPlan(ctx context.Context, claims *auth.Claims, req *bo.CreateTripRequest, pois []*do.PointsOfInterest) (*do.Trip, error) {
	// Get the TripPlanner from the agentHub
	tripPlanner, err := b.agentHub.GetAgent(cnst.AgentTripPlanner)
	if err != nil {
		b.logger.Errorw("failed to get TripPlanner", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	tripPlannerOutput, err := tripPlanner.Execute(ctx, &do.TripPlannerInput{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Duration:  req.Duration,
		POIs:      pois,
	})
	if err != nil {
		b.logger.Errorw("failed to plan trip with TripPlanner", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	output, ok := tripPlannerOutput.(*do.TripPlannerOutput)
	if !ok || output == nil {
		b.logger.Errorw("failed to cast TripPlannerOutput", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	trip := &do.Trip{
		UserID:      claims.UID,
		Title:       output.Title,
		Description: output.Description,
		StartDate:   output.StartDate,
		EndDate:     output.EndDate,
		Days:        int8(output.TotalDays),
	}

	createdTrip, err := b.tripRepo.CreateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to create trip", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	dailyTrips := make([]*do.DailyTrip, 0, len(output.DailyItineraries))
	for _, dailyPlan := range output.DailyItineraries {
		dailyTrip := &do.DailyTrip{
			TripID: createdTrip.ID,
			Day:    int32(dailyPlan.Day),
			Date:   dailyPlan.Date,
			Notes:  dailyPlan.Notes,
		}

		createdDailyTrip, err := b.tripRepo.CreateDailyTrip(ctx, dailyTrip)
		if err != nil {
			b.logger.Errorw("failed to create daily trip", "err", err)
			return nil, xerr.ErrorCreateDailyTripFailed()
		}

		for _, poi := range dailyPlan.POIs {
			dailyItinerary := &do.DailyItinerary{
				TripID:      createdTrip.ID,
				DailyTripID: createdDailyTrip.ID,
				PoiID:       poi.Id,
				Notes:       poi.Notes,
			}

			_, err = b.tripRepo.CreateDailyItinerary(ctx, dailyItinerary)
			if err != nil {
				b.logger.Errorw("failed to create daily itinerary", "err", err)
				return nil, xerr.ErrorCreateDailyItineraryFailed()
			}
		}

		dailyTrips = append(dailyTrips, createdDailyTrip)
	}

	// Get complete trip with all relationships
	completeTrip, err := b.tripRepo.GetTrip(ctx, createdTrip.ID)
	if err != nil {
		b.logger.Errorw("failed to get complete trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	return completeTrip, nil
}

func (b *Trip) CreateTrip(ctx context.Context, req *bo.CreateTripRequest) (*do.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	switch req.CreationMethod {
	case cnst.TripCreationMethodManual:
		// Get the cityPlanner from the agentHub
		cityPlanner, err := b.agentHub.GetAgent(cnst.AgentCityPlanner)
		if err != nil {
			b.logger.Errorw("failed to get CityPlanner", "err", err)
			return nil, xerr.ErrorCreateTripFailed()
		}

		cityPlannerOutput, err := cityPlanner.Execute(ctx, &do.CityPlannerInput{
			Destination: req.Destination,
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
			Duration:    req.Duration,
		})
		if err != nil {
			b.logger.Errorw("failed to plan cities with CityPlanner", "err", err)
			return nil, xerr.ErrorCreateTripFailed()
		}

		rawCities, ok := cityPlannerOutput.(*do.CityPlannerOutput)
		if !ok || len(*rawCities) == 0 {
			b.logger.Errorw("failed to cast CityPlannerOutput", "err", err, "ok", ok, "rawCities", rawCities)
			return nil, xerr.ErrorCreateTripFailed()
		}

		// Get cities IDs from city names
		var cityIds []int32
		for _, dailyCities := range *rawCities {
			for _, cityName := range dailyCities {
				cityId, err := b.cityRepo.GetCityIdByName(ctx, cityName)
				if err != nil {
					b.logger.Errorw("failed to get city id", "err", err)
					return nil, xerr.ErrorCreateTripFailed()
				}
				cityIds = append(cityIds, cityId)
			}
		}

		// Get top POIs for each city
		pois, err := b.poiRepo.GetTopPOIsByCity(ctx, cityIds, 30)
		if err != nil {
			b.logger.Errorw("failed to get top POIs by city", "err", err)
			return nil, xerr.ErrorCreateTripFailed()
		}

		return b.createTripFromPlan(ctx, claims, req, pois)

	case cnst.TripCreationMethodCard:
		poiIds := make([]uuid.UUID, 0, len(req.PoiIds))
		for _, poiId := range req.PoiIds {
			parsedPoiId, err := uuid.Parse(poiId)
			if err != nil {
				continue
			}
			poiIds = append(poiIds, parsedPoiId)
		}
		poiIds = lol.UniqSlice(poiIds)

		var pois []*do.PointsOfInterest
		if len(poiIds) == 0 {
			// get POIs by city IDs
			pois, err = b.poiRepo.GetTopPOIsByCity(ctx, req.CityIds, 30)
			if err != nil {
				b.logger.Errorw("failed to get top POIs by city", "err", err)
				return nil, xerr.ErrorCreateTripFailed()
			}
		} else {
			// get POIs by IDs
			pois, err = b.poiRepo.GetByIds(ctx, poiIds)
			if err != nil {
				b.logger.Errorw("failed to get POIs by ids", "err", err)
				return nil, xerr.ErrorCreateTripFailed()
			}
		}

		return b.createTripFromPlan(ctx, claims, req, pois)

	case cnst.TripCreationMethodExternalLink:
		// TODO: Handle external link creation
		return nil, xerr.ErrorCreateTripFailed()
	case cnst.TripCreationMethodImage:
		// TODO: Handle image-based creation
		return nil, xerr.ErrorCreateTripFailed()
	case cnst.TripCreationMethodVoice:
		// TODO: Handle voice-based creation
		return nil, xerr.ErrorCreateTripFailed()
	default:
		return nil, xerr.ErrorInvalidCreationMethod()
	}
}

func (b *Trip) GetTrip(ctx context.Context, id uuid.UUID) (*do.Trip, error) {
	trip, err := b.tripRepo.GetTrip(ctx, id)
	if err != nil {
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}
	return trip, nil
}

func (b *Trip) UpdateTrip(ctx context.Context, req *bo.UpdateTripRequest) (*do.Trip, error) {
	tripId, err := uuid.Parse(req.ID)
	if err != nil {
		b.logger.Errorw("failed to parse tripRepo id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	trip, err := b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	trip.Title = req.Title
	trip.StartDate = req.StartDate
	trip.EndDate = req.EndDate
	trip.UpdatedAt = time.Now()

	updatedTrip, err := b.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to update tripRepo", "err", err)
		return nil, xerr.ErrorUpdateTripFailed()
	}

	return updatedTrip, nil
}

func (b *Trip) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	if err := b.tripRepo.DeleteTrip(ctx, id); err != nil {
		b.logger.Errorw("failed to delete tripRepo", "err", err)
		return xerr.ErrorDeleteTripFailed()
	}
	return nil
}

func (b *Trip) ListTrips(ctx context.Context) ([]*do.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trips, err := b.tripRepo.ListTrips(ctx, claims.UID)
	if err != nil {
		b.logger.Errorw("failed to list trips", "err", err)
		return nil, xerr.ErrorGetTripListFailed()
	}
	return trips, nil
}

func (b *Trip) CreateDailyTrip(ctx context.Context, req *bo.CreateDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		b.logger.Errorw("failed to parse tripRepo id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	_, err = b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	dailyTrip := &do.DailyTrip{
		TripID: tripId,
		Day:    req.Day,
		Date:   req.Date,
		Notes:  req.Notes,
	}

	createdDailyTrip, err := b.tripRepo.CreateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to create daily tripRepo", "err", err)
		return nil, xerr.ErrorCreateDailyTripFailed()
	}

	return createdDailyTrip, nil
}

func (b *Trip) GetDailyTrip(ctx context.Context, req *bo.GetDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		b.logger.Errorw("failed to get daily tripRepo", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}
	return dailyTrip, nil
}

func (b *Trip) UpdateDailyTrip(ctx context.Context, req *bo.UpdateDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorDailyTripNotFound()
		}
		b.logger.Errorw("daily tripRepo not found", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}

	dailyTrip.Day = req.Day
	dailyTrip.Date = req.Date
	dailyTrip.Notes = req.Notes
	dailyTrip.UpdatedAt = time.Now()

	updatedDailyTrip, err := b.tripRepo.UpdateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to update daily tripRepo", "err", err)
		return nil, xerr.ErrorUpdateDailyTripFailed()
	}

	return updatedDailyTrip, nil
}

func (b *Trip) DeleteDailyTrip(ctx context.Context, req *bo.DeleteDailyTripRequest) error {
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return xerr.ErrorInvalidDailyTripId()
	}

	if err := b.tripRepo.DeleteDailyTrip(ctx, dailyId); err != nil {
		b.logger.Errorw("failed to delete daily tripRepo", "err", err)
		return xerr.ErrorDeleteDailyTripFailed()
	}
	return nil
}

func (b *Trip) ListDailyTrips(ctx context.Context, req *bo.ListDailyTripsRequest) ([]*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}

	dailyTrips, err := b.tripRepo.ListDailyTrips(ctx, tripId)
	if err != nil {
		b.logger.Errorw("failed to list daily trips", "err", err)
		return nil, xerr.ErrorGetDailyTripListFailed()
	}

	return dailyTrips, nil
}

func (b *Trip) ListTripCollaborators(ctx context.Context, tripId uuid.UUID) ([]*do.User, error) {
	// First check if the trip exists
	_, err := b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	// Get collaborators
	collaborators, err := b.tripRepo.ListTripCollaborators(ctx, tripId)
	if err != nil {
		b.logger.Errorw("failed to list trip collaborators", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	return collaborators, nil
}
