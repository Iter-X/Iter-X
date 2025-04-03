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
	StateID uint
	*Pagination
}
