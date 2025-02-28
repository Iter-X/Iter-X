package repository

type BaseRepo[T, R any] interface {
	ToEntity(T) R
	ToEntities([]T) []R
}
