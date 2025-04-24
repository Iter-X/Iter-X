package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/language"
	"github.com/iter-x/iter-x/internal/data/impl/build"
	"go.uber.org/zap"
)

type languageRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func NewLanguage(d *data.Data, logger *zap.SugaredLogger) repository.LanguageRepo {
	return &languageRepositoryImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.language"),
	}
}

func (r *languageRepositoryImpl) ToEntity(po *ent.Language) *do.Language {
	if po == nil {
		return nil
	}
	return build.LanguageRepositoryImplToEntity(po)
}

func (r *languageRepositoryImpl) ToEntities(pos []*ent.Language) []*do.Language {
	if pos == nil {
		return nil
	}
	result := make([]*do.Language, len(pos))
	for i, po := range pos {
		result[i] = r.ToEntity(po)
	}
	return result
}

func (r *languageRepositoryImpl) ListLanguages(ctx context.Context) ([]*do.Language, error) {
	cli := r.GetTx(ctx).Language
	languages, err := cli.Query().Where().All(ctx)
	if err != nil {
		r.logger.Errorw("failed to list languages", "err", err)
		return nil, xerr.ErrorInternalServerError()
	}

	return r.ToEntities(languages), nil
}

func (r *languageRepositoryImpl) FindLanguageByCode(ctx context.Context, code string) (*do.Language, error) {
	cli := r.GetTx(ctx).Language
	lang, err := cli.Query().Where(language.CodeEQ(code)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorInvalidLanguage()
		}
		r.logger.Errorw("failed to find language by code", "code", code, "err", err)
		return nil, xerr.ErrorInternalServerError()
	}

	return r.ToEntity(lang), nil
}
