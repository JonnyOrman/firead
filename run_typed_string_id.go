package firead

func RunTypeStringId[TDocument any]() {
	idReader := StringIdReader{}

	configuration := GenerateConfiguration()

	stringIdDataRetriever := NewStringIdDataRetriever(configuration)

	RunTyped[TDocument, string](
		idReader,
		stringIdDataRetriever)
}
