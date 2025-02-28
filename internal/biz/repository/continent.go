package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Continent[T *ent.Continent, R *do.Continent] interface {
	BaseRepo[T, R]
}

type ContinentRepo = Continent[*ent.Continent, *do.Continent]
