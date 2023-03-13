package stringid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedStringId[TDocument any]() {
	idReader := StringIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewStringIdSnapshotRetriever(configuration)

	firead.RunTyped[TDocument, string](
		idReader,
		snapshotRetriever,
		&configuration,
	)
}
