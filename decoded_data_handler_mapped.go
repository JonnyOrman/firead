package firead

type DecodedDataHandlerMapped[TDocument any, TResponse any] struct {
	documentToResponseMapper DataMapper[TDocument, TResponse]
}

func NewDecodedDataHandlerMapped[TDocument any, TResponse any](
	documentToResponseMapper DataMapper[TDocument, TResponse],
) *DecodedDataHandlerMapped[TDocument, TResponse] {
	this := new(DecodedDataHandlerMapped[TDocument, TResponse])

	this.documentToResponseMapper = documentToResponseMapper

	return this
}

func (this DecodedDataHandlerMapped[TDocument, TResponse]) Handle(decodedData TDocument) (int, any) {
	responseData := this.documentToResponseMapper.Map(decodedData)

	return 200, responseData
}
