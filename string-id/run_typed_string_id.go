package stringid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedStringId[TDocument any]() {
	idReader := NewStringIdReader()

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := firead.NewIdSnapshotRetriever()

	firead.RunTyped[TDocument, string](
		idReader,
		snapshotRetriever,
		&configuration,
	)
}
