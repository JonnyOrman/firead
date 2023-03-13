package intid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedIntId[TDocument any]() {
	idReader := IntIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewIntIdSnapshotRetriever(configuration)

	dataRetriever := firead.NewFirestoreDataRetriever[int](configuration, snapshotRetriever)

	firead.RunTyped[TDocument, int](
		idReader,
		dataRetriever)
}
