package repository

import (
	"context"

	"github.com/google/uuid"
	
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type PoiFiles[T *ent.PointsOfInterestFiles, R *do.PointsOfInterestFiles] interface {
	BaseRepo[T, R]

	FindByPoiID(ctx context.Context, poiID uuid.UUID) ([]R, error)
}

type PoiFilesRepo = PoiFiles[*ent.PointsOfInterestFiles, *do.PointsOfInterestFiles]
