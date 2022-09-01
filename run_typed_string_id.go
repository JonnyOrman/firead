package firead

import "github.com/jonnyorman/fireworks"

func RunTypeStringId[TDocument any]() {
	idReader := StringIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	stringIdDataRetriever := NewStringIdDataRetriever(configuration)

	RunTyped[TDocument, string](
		idReader,
		stringIdDataRetriever)
}
