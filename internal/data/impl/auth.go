package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/cache"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/data/ent/user"
	"github.com/iter-x/iter-x/internal/data/ent/userpreference"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewAuth(c *conf.Auth, d *data.Data, logger *zap.SugaredLogger) repository.AuthRepo {
	return &authRepositoryImpl{
		smsConf: c.GetSmsCode(),
		Tx:      d.Tx,
		Cacher:  d.Cache,
		logger:  logger.Named("repo.auth"),
	}
}

func NewUser(d *data.Data, logger *zap.SugaredLogger) repository.UserRepo {
	return &authRepositoryImpl{
		Tx:     d.Tx,
		Cacher: d.Cache,
		logger: logger.Named("repo.user"),
	}
}

type authRepositoryImpl struct {
	smsConf *conf.Auth_SmsCode
	*data.Tx
	cache.Cacher
	logger *zap.SugaredLogger
}

func (r *authRepositoryImpl) ToEntity(po *ent.User) *do.User {
	if po == nil {
		return nil
	}
	return build.AuthRepositoryImplToEntity(po)
}

func (r *authRepositoryImpl) ToEntities(pos []*ent.User) []*do.User {
	if pos == nil {
		return nil
	}
	return build.AuthRepositoryImplToEntities(pos)
}

func (r *authRepositoryImpl) FindByEmail(ctx context.Context, email string) (*do.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error) {
	cli := r.GetTx(ctx).User
	usr, err := cli.Get(ctx, id)
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) Create(ctx context.Context, user *do.User) (*do.User, error) {
	cli := r.GetTx(ctx).User

	row, err := cli.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		SetNickname(user.Nickname).
		SetRemark(user.Remark).
		SetPhone(user.Phone).
		SetAvatarURL(user.AvatarURL).
		SetEmail(user.Email).
		SetStatus(user.Status.GetValue()).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *authRepositoryImpl) Update(ctx context.Context, user *do.User) (*do.User, error) {
	cli := r.GetTx(ctx).User

	row, err := cli.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *authRepositoryImpl) GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	row, err := cli.Query().Where(refreshtoken.UserID(userId)).Only(ctx)
	return build.RefreshTokenImplToEntity(row), err
}

func (r *authRepositoryImpl) GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	row, err := cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
	return build.RefreshTokenImplToEntity(row), err
}

func (r *authRepositoryImpl) SaveRefreshToken(ctx context.Context, val *do.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.Create().
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUserID(val.UserID).
		Save(ctx)
	return err
}

func (r *authRepositoryImpl) UpdateRefreshToken(ctx context.Context, val *do.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.UpdateOneID(val.ID).
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUpdatedAt(val.UpdatedAt).
		Save(ctx)
	return err
}

func (r *authRepositoryImpl) FindByPhone(ctx context.Context, phone string) (*do.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Query().Where(user.PhoneEQ(phone)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) GetUserPreference(ctx context.Context, userId uuid.UUID) (*do.UserPreference, error) {
	cli := r.GetTx(ctx).UserPreference
	userPref, err := cli.Query().Where(userpreference.UserID(userId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
		r.logger.Errorw("failed to find user preference by id", "err", err)
		return nil, xerr.ErrorInternalServerError()
	}

	return build.UserPreferenceRepositoryImplToEntity(userPref), nil
}

func (r *authRepositoryImpl) CreateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error {
	cli := r.GetTx(ctx).UserPreference

	// Check if preference already exists for this user
	exists, err := cli.Query().Where(userpreference.UserID(userId)).Exist(ctx)
	if err != nil {
		r.logger.Errorw("failed to check user preference existence", "err", err)
		return xerr.ErrorInternalServerError()
	}
	if exists {
		r.logger.Errorw("user preference already exists", "userId", userId)
		return xerr.ErrorBadRequest()
	}

	_, err = cli.Create().
		SetUserID(userId).
		SetAppLanguage(pref.AppLanguage).
		SetDefaultCity(pref.DefaultCity).
		SetTimeFormat(userpreference.TimeFormat(pref.TimeFormat)).
		SetDistanceUnit(userpreference.DistanceUnit(pref.DistanceUnit)).
		SetDarkMode(userpreference.DarkMode(pref.DarkMode)).
		SetNotifyItinerary(pref.TripReminder).
		SetNotifyCommunity(pref.CommunityNotification).
		SetNotifyRecommendations(pref.RecommendContentPush).
		Save(ctx)
	if err != nil {
		r.logger.Errorw("failed to create user preference", "err", err)
		return xerr.ErrorInternalServerError()
	}

	return nil
}

func (r *authRepositoryImpl) UpdateUserPreference(ctx context.Context, userId uuid.UUID, pref *do.UserPreference) error {
	cli := r.GetTx(ctx).UserPreference

	// Check if the record exists first
	userPref, err := cli.Query().Where(userpreference.UserID(userId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			r.logger.Errorw("user preference not found", "userId", userId)
			return xerr.ErrorBadRequest()
		}
		r.logger.Errorw("failed to check user preference existence", "err", err)
		return xerr.ErrorInternalServerError()
	}

	_, err = cli.UpdateOneID(userPref.ID).
		SetAppLanguage(pref.AppLanguage).
		SetDefaultCity(pref.DefaultCity).
		SetTimeFormat(userpreference.TimeFormat(pref.TimeFormat)).
		SetDistanceUnit(userpreference.DistanceUnit(pref.DistanceUnit)).
		SetDarkMode(userpreference.DarkMode(pref.DarkMode)).
		SetNotifyItinerary(pref.TripReminder).
		SetNotifyCommunity(pref.CommunityNotification).
		SetNotifyRecommendations(pref.RecommendContentPush).
		Save(ctx)
	if err != nil {
		r.logger.Errorw("failed to update user preference", "err", err)
		return xerr.ErrorInternalServerError()
	}

	return nil
}
