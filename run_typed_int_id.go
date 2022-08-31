package firead

func RunTypeIntId[TDocument any]() {
	idReader := IntIdReader{}

	configuration := GenerateConfiguration()

	intIdDataRetriever := NewIntIdDataRetriever(configuration)

	RunTyped[TDocument, int](
		idReader,
		intIdDataRetriever)
}
