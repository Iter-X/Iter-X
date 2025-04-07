package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func UserPreferenceRepositoryImplToEntity(po *ent.UserPreference) *do.UserPreference {
	if po == nil {
		return nil
	}

	return &do.UserPreference{
		ID:                    po.ID,
		UserID:                po.UserID,
		AppLanguage:           po.AppLanguage,
		DefaultCity:           po.DefaultCity,
		TimeFormat:            string(po.TimeFormat),
		DistanceUnit:          string(po.DistanceUnit),
		DarkMode:              string(po.DarkMode),
		TripReminder:          po.NotifyItinerary,
		CommunityNotification: po.NotifyCommunity,
		RecommendContentPush:  po.NotifyRecommendations,
		CreatedAt:             po.CreatedAt,
		UpdatedAt:             po.UpdatedAt,
	}
}
