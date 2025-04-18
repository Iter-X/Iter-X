package cnst

type TripCreationMethod int

const (
	// TripCreationMethodUnspecified represents an unspecified creation method
	TripCreationMethodUnspecified TripCreationMethod = iota // unspecified
	// TripCreationMethodManual represents manual trip creation
	TripCreationMethodManual // manual
	// TripCreationMethodCard represents card-based trip creation
	TripCreationMethodCard // card
	// TripCreationMethodExternalLink represents external link based trip creation
	TripCreationMethodExternalLink // external_link
	// TripCreationMethodImage represents image-based trip creation
	TripCreationMethodImage // image
	// TripCreationMethodVoice represents voice-based trip creation
	TripCreationMethodVoice // voice
)

const MaxDay = 99
