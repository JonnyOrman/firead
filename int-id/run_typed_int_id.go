package intid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedIntId[TDocument any]() {
	idReader := NewIntIdReader()

	configuration := fireworks.GenerateConfiguration("firead-config")

	stringSnapshotRetriever := firead.NewIdSnapshotRetriever()

	snapshotRetriever := NewIntIdSnapshotRetriever(stringSnapshotRetriever)

	firead.RunTyped[TDocument, int](
		idReader,
		snapshotRetriever,
		&configuration,
	)
}
