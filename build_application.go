package firead

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

func BuildApplication[TDocument any, TId Id](
	idReader ParamReader[TId],
	snapshotRetriever SnapshotRetriever[TId],
	configuration *fireworks.Configuration,
) *fireworks.Application {
	dataDecoder := NewMapStructureDataDecoder[TDocument]()

	decodedDataHandler := NewDecodedDataHandler[TDocument]()

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
