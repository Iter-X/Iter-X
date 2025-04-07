package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Language[T *ent.Language, R *do.Language] interface {
	BaseRepo[T, R]

	ListLanguages(ctx context.Context) ([]*do.Language, error)
	FindLanguageByCode(ctx context.Context, code string) (*do.Language, error)
}

type LanguageRepo = Language[*ent.Language, *do.Language]
