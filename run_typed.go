package firead

func RunTyped[TDocument any, TId Id](
	idReader ParamReader[TId],
	dataRetriever DataRetriever[TId]) {
	application := BuildApplication[TDocument](
		idReader,
		dataRetriever)

	application.Run()
}
