package intid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunMappedIntId[TDocument any, TResponse any](dataMapper firead.DataMapper[TDocument, TResponse]) {
	idReader := NewIntIdReader()

	configuration := fireworks.GenerateConfiguration("firead-config")

	stringSnapshotRetriever := firead.NewIdSnapshotRetriever()

	snapshotRetriever := NewIntIdSnapshotRetriever(stringSnapshotRetriever)

	firead.RunMapped[TDocument, TResponse, int](
		idReader,
		snapshotRetriever,
		dataMapper,
		&configuration,
	)
}
