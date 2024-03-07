package firead

import "github.com/jonnyorman/fireworks"

func RunMapped[TDocument any, TResponse any, TId Id](
	idReader ParamReader[TId],
	snapshotRetriever SnapshotRetriever[TId],
	documentToResponseMapper DataMapper[TDocument, TResponse],
	configuration *fireworks.Configuration,
) {
	application := BuildApplicationMapped[TDocument, TResponse](
		idReader,
		snapshotRetriever,
		documentToResponseMapper,
		configuration,
	)

	application.Run()
}
