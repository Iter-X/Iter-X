package bo

import (
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
)

type SearchPointsOfInterestParams struct {
	Keyword        string `json:"keyword"`
	Limit          int    `json:"limit"`
	GeographyLevel poiV1.SearchPointsOfInterestRequest_GeographyLevel
	depth          uint8
}

// DepthDec Search depth -1
func (s *SearchPointsOfInterestParams) DepthDec() *SearchPointsOfInterestParams {
	s.depth--
	switch s.GeographyLevel {
	case poiV1.SearchPointsOfInterestRequest_POI:
		s.GeographyLevel = poiV1.SearchPointsOfInterestRequest_CITY
	case poiV1.SearchPointsOfInterestRequest_CITY:
		s.GeographyLevel = poiV1.SearchPointsOfInterestRequest_STATE
	case poiV1.SearchPointsOfInterestRequest_STATE:
		s.GeographyLevel = poiV1.SearchPointsOfInterestRequest_COUNTRY
	case poiV1.SearchPointsOfInterestRequest_COUNTRY:
		s.GeographyLevel = poiV1.SearchPointsOfInterestRequest_CONTINENT
	default:
		s.GeographyLevel = poiV1.SearchPointsOfInterestRequest_POI
	}
	return s
}

// WithDepth set search depth
func (s *SearchPointsOfInterestParams) WithDepth(depth uint8) *SearchPointsOfInterestParams {
	s.depth = depth
	return s
}

// Depth get search depth
func (s *SearchPointsOfInterestParams) Depth() uint8 {
	return s.depth
}

// IsNext if has next
func (s *SearchPointsOfInterestParams) IsNext() bool {
	return s.depth > 0
}

// IsPoi is poi
func (s *SearchPointsOfInterestParams) IsPoi() bool {
	return s.GeographyLevel == poiV1.SearchPointsOfInterestRequest_POI
}

// IsCity is city
func (s *SearchPointsOfInterestParams) IsCity() bool {
	return s.GeographyLevel == poiV1.SearchPointsOfInterestRequest_CITY
}

// IsState is state
func (s *SearchPointsOfInterestParams) IsState() bool {
	return s.GeographyLevel == poiV1.SearchPointsOfInterestRequest_STATE
}

// IsCountry is country
func (s *SearchPointsOfInterestParams) IsCountry() bool {
	return s.GeographyLevel == poiV1.SearchPointsOfInterestRequest_COUNTRY
}

// IsContinent is continent
func (s *SearchPointsOfInterestParams) IsContinent() bool {
	return s.GeographyLevel == poiV1.SearchPointsOfInterestRequest_CONTINENT
}
