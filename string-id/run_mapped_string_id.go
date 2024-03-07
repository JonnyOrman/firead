package stringid

import (
	"github.com/jonnyorman/firead"
	"github.com/jonnyorman/fireworks"
)

func RunMappedStringId[TDocument any, TResponse any](dataMapper firead.DataMapper[TDocument, TResponse]) {
	idReader := NewStringIdReader()

	configuration := fireworks.GenerateConfiguration("firead-config")

	snapshotRetriever := firead.NewIdSnapshotRetriever()

	firead.RunMapped[TDocument, TResponse, string](
		idReader,
		snapshotRetriever,
		dataMapper,
		&configuration,
	)
}
