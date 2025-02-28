package repository

type Base[T, R any] interface {
	ToEntity(T) R
	ToEntities([]T) []R
}
