package firead

type DecodedDataHandler[T any] struct{}

func NewDecodedDataHandler[T any]() *DecodedDataHandler[T] {
	this := new(DecodedDataHandler[T])

	return this
}

func (this DecodedDataHandler[T]) Handle(decodedData T) (int, any) {
	return 200, decodedData
}
