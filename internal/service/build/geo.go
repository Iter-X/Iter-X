package build

import (
	"context"

	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/pkg/util"
)

// ToContinentProto converts a domain continent object to a protobuf continent
func ToContinentProto(ctx context.Context, continent *do.Continent) *geoV1.Continent {
	if continent == nil {
		return nil
	}

	return &geoV1.Continent{
		Id:        uint32(continent.ID),
		Name:      util.GetLocalizedName(ctx, continent.NameEn, continent.NameCn),
		NameLocal: continent.NameLocal,
		NameEn:    continent.NameEn,
		NameCn:    continent.NameCn,
		Code:      continent.Code,
	}
}

// ToContinentsProto converts a slice of domain continent objects to protobuf continents
func ToContinentsProto(ctx context.Context, continents []*do.Continent) []*geoV1.Continent {
	if continents == nil {
		return nil
	}
	var result []*geoV1.Continent
	for _, continent := range continents {
		result = append(result, ToContinentProto(ctx, continent))
	}
	return result
}

// ToCountryProto converts a domain country object to a protobuf country
func ToCountryProto(ctx context.Context, country *do.Country) *geoV1.Country {
	if country == nil {
		return nil
	}

	countryProto := &geoV1.Country{
		Id:          uint32(country.ID),
		Name:        util.GetLocalizedName(ctx, country.NameEn, country.NameCn),
		NameLocal:   country.NameLocal,
		NameEn:      country.NameEn,
		NameCn:      country.NameCn,
		Code:        country.Code,
		ContinentId: uint32(country.ContinentID),
		ImageUrl:    country.ImageUrl,
	}

	if country.Continent != nil {
		countryProto.Continent = ToContinentProto(ctx, country.Continent)
	}

	return countryProto
}

// ToCountriesProto converts a slice of domain country objects to protobuf countries
func ToCountriesProto(ctx context.Context, countries []*do.Country) []*geoV1.Country {
	if countries == nil {
		return nil
	}
	var result []*geoV1.Country
	for _, country := range countries {
		result = append(result, ToCountryProto(ctx, country))
	}
	return result
}

// ToStateProto converts a domain state object to a protobuf state
func ToStateProto(ctx context.Context, state *do.State) *geoV1.State {
	if state == nil {
		return nil
	}

	stateProto := &geoV1.State{
		Id:        uint32(state.ID),
		Name:      util.GetLocalizedName(ctx, state.NameEn, state.NameCn),
		NameLocal: state.NameLocal,
		NameEn:    state.NameEn,
		NameCn:    state.NameCn,
		Code:      state.Code,
		CountryId: uint32(state.CountryID),
	}

	if state.Country != nil {
		stateProto.Country = ToCountryProto(ctx, state.Country)
	}

	return stateProto
}

// ToStatesProto converts a slice of domain state objects to protobuf states
func ToStatesProto(ctx context.Context, states []*do.State) []*geoV1.State {
	if states == nil {
		return nil
	}
	var result []*geoV1.State
	for _, state := range states {
		result = append(result, ToStateProto(ctx, state))
	}
	return result
}

// ToCityProto converts a domain city object to a protobuf city
func ToCityProto(ctx context.Context, city *do.City) *geoV1.City {
	if city == nil {
		return nil
	}

	cityProto := &geoV1.City{
		Id:        uint32(city.ID),
		Name:      util.GetLocalizedName(ctx, city.NameEn, city.NameCn),
		NameEn:    city.NameEn,
		NameCn:    city.NameCn,
		NameLocal: city.NameLocal,
		Code:      city.Code,
		StateId:   uint32(city.StateID),
	}

	if city.State != nil {
		cityProto.State = ToStateProto(ctx, city.State)
	}

	return cityProto
}

// ToCitiesProto converts a slice of domain city objects to protobuf cities
func ToCitiesProto(ctx context.Context, cities []*do.City) []*geoV1.City {
	if cities == nil {
		return nil
	}
	var result []*geoV1.City
	for _, city := range cities {
		result = append(result, ToCityProto(ctx, city))
	}
	return result
}

// ToPOIProto converts a domain POI object to a protobuf POI
func ToPOIProto(ctx context.Context, poi *do.PointsOfInterest) *geoV1.POI {
	if poi == nil {
		return nil
	}

	poiProto := &geoV1.POI{
		Id:           poi.ID.String(),
		Name:         util.GetLocalizedName(ctx, poi.NameEn, poi.NameCn),
		NameLocal:    poi.NameLocal,
		NameEn:       poi.NameEn,
		NameCn:       poi.NameCn,
		Description:  poi.Description,
		ImageUrl:     poi.ImageUrl,
		Rating:       float64(poi.Rating),
		ReviewsCount: uint32(len(poi.PoiFiles)),
		CityId:       uint32(poi.CityID),
		Latitude:     poi.Latitude,
		Longitude:    poi.Longitude,
		Duration:     poi.RecommendedDurationMinutes,
		Popularity:   50, // TODO: Add popularity calculation
	}

	if poi.City != nil {
		poiProto.City = ToCityProto(ctx, poi.City)
	}

	return poiProto
}

// ToPOIsProto converts a slice of domain POI objects to protobuf POIs
func ToPOIsProto(ctx context.Context, pois []*do.PointsOfInterest) []*geoV1.POI {
	if pois == nil {
		return nil
	}
	var result []*geoV1.POI
	for _, poi := range pois {
		result = append(result, ToPOIProto(ctx, poi))
	}
	return result
}
