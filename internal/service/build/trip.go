package build

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
)

func ToTripProto(t *do.Trip) *tripV1.Trip {
	if t == nil {
		return nil
	}
	return &tripV1.Trip{
		Id:          t.ID.String(),
		Status:      t.Status,
		Title:       t.Title,
		Description: t.Description,
		StartTs:     timestamppb.New(t.StartDate),
		EndTs:       timestamppb.New(t.EndDate),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
		DailyTrips:  ToDailyTripsProto(t.DailyTrip),
		PoiPool:     ToPointOfInterestsProto(t.PoiPool),
	}
}

func ToTripsProto(ts []*do.Trip) []*tripV1.Trip {
	list := make([]*tripV1.Trip, 0, len(ts))
	for _, v := range ts {
		list = append(list, ToTripProto(v))
	}
	return list
}

func ToDailyTripProto(dt *do.DailyTrip) *tripV1.DailyTrip {
	if dt == nil {
		return nil
	}
	return &tripV1.DailyTrip{
		Id:               dt.ID.String(),
		TripId:           dt.TripID.String(),
		Day:              dt.Day,
		Date:             timestamppb.New(dt.Date),
		Notes:            dt.Notes,
		CreatedAt:        timestamppb.New(dt.CreatedAt),
		UpdatedAt:        timestamppb.New(dt.UpdatedAt),
		DailyItineraries: ToDailyItinerariesProto(dt.DailyItinerary),
	}
}

func ToDailyTripsProto(dts []*do.DailyTrip) []*tripV1.DailyTrip {
	list := make([]*tripV1.DailyTrip, 0, len(dts))
	for _, v := range dts {
		list = append(list, ToDailyTripProto(v))
	}
	return list
}

func ToDailyItineraryProto(di *do.DailyItinerary) *tripV1.DailyItinerary {
	if di == nil {
		return nil
	}
	return &tripV1.DailyItinerary{
		Id:          di.ID.String(),
		TripId:      di.TripID.String(),
		DailyTripId: di.DailyTripID.String(),
		PoiId:       di.PoiID.String(),
		Notes:       di.Notes,
		CreatedAt:   timestamppb.New(di.CreatedAt),
		UpdatedAt:   timestamppb.New(di.UpdatedAt),
		Poi:         ToPointOfInterestProto(di.Poi),
	}
}

func ToPointOfInterestProto(poi *do.PointsOfInterest) *poiV1.PointOfInterest {
	if poi == nil {
		return nil
	}
	var city, state, country string
	if poi.City != nil {
		city = poi.City.NameEn
	}
	if poi.State != nil {
		state = poi.State.NameEn
	}
	if poi.Country != nil {
		country = poi.Country.NameEn
	}
	return &poiV1.PointOfInterest{
		Id:                         poi.ID.String(),
		Name:                       poi.NameEn,
		NameEn:                     poi.NameEn,
		NameCn:                     poi.NameCn,
		NameLocal:                  poi.NameLocal,
		Description:                poi.Description,
		Address:                    poi.Address,
		Latitude:                   poi.Latitude,
		Longitude:                  poi.Longitude,
		Type:                       poi.Type,
		Category:                   poi.Category,
		Rating:                     poi.Rating,
		RecommendedDurationMinutes: poi.RecommendedDurationMinutes,
		City:                       city,
		State:                      state,
		Country:                    country,
	}
}

func ToDailyItinerariesProto(dis []*do.DailyItinerary) []*tripV1.DailyItinerary {
	if dis == nil {
		return nil
	}
	list := make([]*tripV1.DailyItinerary, 0, len(dis))
	for _, v := range dis {
		list = append(list, ToDailyItineraryProto(v))
	}
	return list
}

func ToTripCollaboratorProto(collaborator *do.TripCollaborator) *tripV1.TripCollaborator {
	if collaborator == nil {
		return nil
	}
	return &tripV1.TripCollaborator{
		Id:        collaborator.ID.String(),
		Username:  collaborator.Username,
		Nickname:  collaborator.Nickname,
		AvatarUrl: collaborator.AvatarURL,
		Status:    tripV1.CollaboratorStatus(tripV1.CollaboratorStatus_value[collaborator.Status]),
	}
}

func ToTripCollaboratorsProto(collaborators []*do.TripCollaborator) []*tripV1.TripCollaborator {
	if collaborators == nil {
		return nil
	}
	list := make([]*tripV1.TripCollaborator, 0, len(collaborators))
	for _, collaborator := range collaborators {
		list = append(list, ToTripCollaboratorProto(collaborator))
	}
	return list
}

// ToPointOfInterestsProto converts a slice of TripPOIPool to a slice of PointOfInterest
func ToPointOfInterestsProto(poiPools []*do.TripPOIPool) []*poiV1.PointOfInterest {
	if poiPools == nil {
		return nil
	}
	protoPois := make([]*poiV1.PointOfInterest, 0, len(poiPools))
	for _, pool := range poiPools {
		if pool.Poi != nil {
			protoPois = append(protoPois, ToPointOfInterestProto(pool.Poi))
		}
	}
	return protoPois
}
