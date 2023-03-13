package firead

import "github.com/jonnyorman/fireworks"

func RunTyped[TDocument any, TId Id](
	idReader ParamReader[TId],
	snapshotRetriever SnapshotRetriever[TId],
	configuration *fireworks.Configuration,
) {
	application := BuildApplication[TDocument](
		idReader,
		snapshotRetriever,
		configuration,
	)

	application.Run()
}
