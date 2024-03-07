package firead

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

func BuildApplicationMapped[TDocument any, TResponse any, TId Id](
	idReader ParamReader[TId],
	snapshotRetriever SnapshotRetriever[TId],
	documentToResponseMapper DataMapper[TDocument, TResponse],
	configuration *fireworks.Configuration,
) *fireworks.Application {
	dataDecoder := NewMapStructureDataDecoder[TDocument]()

	decodedDataHandler := NewDecodedDataHandlerMapped[TDocument](
		documentToResponseMapper,
	)

	existingDocumentDataHandler := NewExistingDocumentDataHandler[TDocument](
		dataDecoder,
		decodedDataHandler,
	)

	nonExistingDocumentHandler := NewNonExistingDocumentHandler()

	responseResolver := NewDocumentResponseResolver(
		configuration,
		snapshotRetriever,
		existingDocumentDataHandler,
		nonExistingDocumentHandler,
	)

	responseWriter := NewJsonResponseWriter()

	requestHandler := NewDocumentRequestHandler[TDocument, TId](
		idReader,
		responseResolver,
		responseWriter,
	)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/:id", requestHandler.Handle)

	application := fireworks.NewApplication(router)

	return application
}
