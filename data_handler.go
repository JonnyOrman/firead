package firead

type DataHandler[TData any] interface {
	Handle(data TData) (int, any)
}
