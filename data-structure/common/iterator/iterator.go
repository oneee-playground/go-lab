package iterator

type Iterator[T any] interface {
	IsEnd() bool
	Get() T
	Next()
}
