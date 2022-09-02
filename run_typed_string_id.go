package firead

import "github.com/jonnyorman/fireworks"

func RunTypeStringId[TDocument any]() {
	idReader := StringIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewStringIdSnapshotRetriever(configuration)

	dataRetriever := NewFirestoreDataRetriever[string](configuration, snapshotRetriever)

	RunTyped[TDocument, string](
		idReader,
		dataRetriever)
}
