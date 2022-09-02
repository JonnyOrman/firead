package firead

import "github.com/jonnyorman/fireworks"

func RunTypedIntId[TDocument any]() {
	idReader := IntIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewIntIdSnapshotRetriever(configuration)

	dataRetriever := NewFirestoreDataRetriever[int](configuration, snapshotRetriever)

	RunTyped[TDocument, int](
		idReader,
		dataRetriever)
}
