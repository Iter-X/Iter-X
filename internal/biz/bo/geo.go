package bo

// ListContinentsParams parameters for listing continents
type ListContinentsParams struct {
	*Pagination
}

// ListCountriesParams parameters for listing countries
type ListCountriesParams struct {
	// Optional continent ID filter
	ContinentID uint
	*Pagination
}

// ListStatesParams parameters for listing states/provinces
type ListStatesParams struct {
	// Optional country ID filter
	CountryID uint
	*Pagination
}

// ListCitiesParams parameters for listing cities
type ListCitiesParams struct {
	// Optional state ID filter
	StateId *uint32
	// Optional country ID filter
	CountryId *uint32
	*Pagination
}

// ListPOIsParams parameters for listing POIs
type ListPOIsParams struct {
	// Optional city ID filter
	CityId *uint32
	// Optional keyword filter
	Keyword *string
	// Optional city IDs filter
	CityIds []uint32
	*Pagination
}
