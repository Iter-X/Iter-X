package bo

type UpdateUserInfoRequest struct {
	Username  string
	Nickname  string
	AvatarURL string
}

type UserPreference struct {
	AppLanguage           string
	DefaultCity           string
	TimeFormat            string
	DistanceUnit          string
	DarkMode              string
	TripReminder          bool
	CommunityNotification bool
	ContentPush           bool
}
