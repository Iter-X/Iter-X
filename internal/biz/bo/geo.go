package bo

// ListContinentsParams parameters for listing continents
type ListContinentsParams struct {
	// Number of records per page
	Limit int32
	// Offset, starting from 0
	Offset int
}

// ListCountriesParams parameters for listing countries
type ListCountriesParams struct {
	// Optional continent ID filter
	ContinentID uint
	// Number of records per page
	Limit int32
	// Offset, starting from 0
	Offset int
}

// ListStatesParams parameters for listing states/provinces
type ListStatesParams struct {
	// Optional country ID filter
	CountryID uint
	// Number of records per page
	Limit int32
	// Offset, starting from 0
	Offset int
}

// ListCitiesParams parameters for listing cities
type ListCitiesParams struct {
	// Optional state ID filter
	StateID uint
	// Number of records per page
	Limit int32
	// Offset, starting from 0
	Offset int
}
