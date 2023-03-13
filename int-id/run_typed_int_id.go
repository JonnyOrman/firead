package intid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunTypedIntId[TDocument any]() {
	idReader := IntIdReader{}

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := NewIntIdSnapshotRetriever(configuration)

	firead.RunTyped[TDocument, int](
		idReader,
		snapshotRetriever,
		&configuration,
	)
}
