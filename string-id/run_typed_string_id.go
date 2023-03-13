package stringid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedStringId[TDocument any]() {
	idReader := StringIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewStringIdSnapshotRetriever(configuration)

	dataRetriever := firead.NewFirestoreDataRetriever[string](configuration, snapshotRetriever)

	firead.RunTyped[TDocument, string](
		idReader,
		dataRetriever)
}
