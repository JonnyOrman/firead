package firead

import "github.com/jonnyorman/fireworks"

func RunTypeIntId[TDocument any]() {
	idReader := IntIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	intIdDataRetriever := NewIntIdDataRetriever(configuration)

	RunTyped[TDocument, int](
		idReader,
		intIdDataRetriever)
}
