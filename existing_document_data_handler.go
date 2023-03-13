package firead

type ExistingDocumentDataHandler[T any] struct {
	dataDecoder        DataDecoder[T]
	decodedDataHandler DataHandler[T]
}

func NewExistingDocumentDataHandler[T any](
	dataDecoder DataDecoder[T],
	decodedDataHandler DataHandler[T],
) *ExistingDocumentDataHandler[T] {
	this := new(ExistingDocumentDataHandler[T])

	this.dataDecoder = dataDecoder
	this.decodedDataHandler = decodedDataHandler

	return this
}

func (this ExistingDocumentDataHandler[T]) Handle(data map[string]interface{}) (int, any) {
	decodedData := this.dataDecoder.Decode(data)

	return this.decodedDataHandler.Handle(decodedData)
}
