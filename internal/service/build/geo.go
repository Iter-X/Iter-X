package build

import (
	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
)

// ToContinentProto converts a domain continent object to a protobuf continent
func ToContinentProto(continent *do.Continent) *geoV1.Continent {
	if continent == nil {
		return nil
	}
	return &geoV1.Continent{
		Id:     uint32(continent.ID),
		Name:   continent.Name,
		NameEn: continent.NameEn,
		NameCn: continent.NameCn,
		Code:   continent.Code,
	}
}

// ToContinentsProto converts a slice of domain continent objects to protobuf continents
func ToContinentsProto(continents []*do.Continent) []*geoV1.Continent {
	if continents == nil {
		return nil
	}
	var result []*geoV1.Continent
	for _, continent := range continents {
		result = append(result, ToContinentProto(continent))
	}
	return result
}

// ToCountryProto converts a domain country object to a protobuf country
func ToCountryProto(country *do.Country) *geoV1.Country {
	if country == nil {
		return nil
	}
	countryProto := &geoV1.Country{
		Id:          uint32(country.ID),
		Name:        country.Name,
		NameEn:      country.NameEn,
		NameCn:      country.NameCn,
		Code:        country.Code,
		ContinentId: uint32(country.ContinentID),
	}

	if country.Continent != nil {
		countryProto.Continent = ToContinentProto(country.Continent)
	}

	return countryProto
}

// ToCountriesProto converts a slice of domain country objects to protobuf countries
func ToCountriesProto(countries []*do.Country) []*geoV1.Country {
	if countries == nil {
		return nil
	}
	var result []*geoV1.Country
	for _, country := range countries {
		result = append(result, ToCountryProto(country))
	}
	return result
}

// ToStateProto converts a domain state object to a protobuf state
func ToStateProto(state *do.State) *geoV1.State {
	if state == nil {
		return nil
	}
	stateProto := &geoV1.State{
		Id:        uint32(state.ID),
		Name:      state.Name,
		NameEn:    state.NameEn,
		NameCn:    state.NameCn,
		Code:      state.Code,
		CountryId: uint32(state.CountryID),
	}

	if state.Country != nil {
		stateProto.Country = ToCountryProto(state.Country)
	}

	return stateProto
}

// ToStatesProto converts a slice of domain state objects to protobuf states
func ToStatesProto(states []*do.State) []*geoV1.State {
	if states == nil {
		return nil
	}
	var result []*geoV1.State
	for _, state := range states {
		result = append(result, ToStateProto(state))
	}
	return result
}

// ToCityProto converts a domain city object to a protobuf city
func ToCityProto(city *do.City) *geoV1.City {
	if city == nil {
		return nil
	}
	cityProto := &geoV1.City{
		Id:      uint32(city.ID),
		Name:    city.Name,
		NameEn:  city.NameEn,
		NameCn:  city.NameCn,
		Code:    city.Code,
		StateId: uint32(city.StateID),
	}

	if city.State != nil {
		cityProto.State = ToStateProto(city.State)
	}

	return cityProto
}

// ToCitiesProto converts a slice of domain city objects to protobuf cities
func ToCitiesProto(cities []*do.City) []*geoV1.City {
	if cities == nil {
		return nil
	}
	var result []*geoV1.City
	for _, city := range cities {
		result = append(result, ToCityProto(city))
	}
	return result
}
