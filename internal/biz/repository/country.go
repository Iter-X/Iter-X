package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Country[T *ent.Country, R *do.Country] interface {
	BaseRepo[T, R]
}

type CountryRepo = Country[*ent.Country, *do.Country]
