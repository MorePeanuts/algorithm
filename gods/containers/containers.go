package containers

type Container[T any] interface {
	IsEmpty() bool
	Len() int
	Clear()
}
