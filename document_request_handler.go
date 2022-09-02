package firead

import "github.com/gin-gonic/gin"

type DocumentRequestHandler[TDocument any, TId any] struct {
	paramReader    ParamReader[TId]
	documentReader DocumentReader[TDocument, TId]
	responseWriter ResponseWriter
}

func NewDocumentRequestHandler[TDocument any, TId any](

	paramReader ParamReader[TId],
	documentReader DocumentReader[TDocument, TId],
	responseSetter ResponseWriter,
) *DocumentRequestHandler[TDocument, TId] {
	this := new(DocumentRequestHandler[TDocument, TId])

	this.paramReader = paramReader
	this.documentReader = documentReader
	this.responseWriter = responseSetter

	return this
}

func (this DocumentRequestHandler[TDocument, TId]) Handle(ginContext *gin.Context) {
	id := this.paramReader.Read(ginContext)

	document := this.documentReader.Read(id)

	//ginContext.JSON(200, document)
	this.responseWriter.Write(ginContext, 200, document)
}
