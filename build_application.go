package firead

import "github.com/jonnyorman/fireworks"

func BuildApplication[TDocument any, TId Id](
	idReader ParamReader[TId],
	dataRetriever DataRetriever[TId]) *fireworks.Application {
	documentReader := NewFirestoreDocumentReader[TDocument](dataRetriever)

	responseWriter := NewJsonResponseWriter()

	requestHandler := NewDocumentRequestHandler[TDocument, TId](
		idReader,
		documentReader,
		responseWriter)

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddGet("/:id", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
