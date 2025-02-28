package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type City[T *ent.City, R *do.City] interface {
	BaseRepo[T, R]
}

type CityRepo = City[*ent.City, *do.City]
