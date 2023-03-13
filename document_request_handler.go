package firead

import "github.com/gin-gonic/gin"

type DocumentRequestHandler[TDocument any, TId Id] struct {
	paramReader      ParamReader[TId]
	responseResolver ResponseResolver[TId]
	responseWriter   ResponseWriter
}

func NewDocumentRequestHandler[TDocument any, TId Id](

	paramReader ParamReader[TId],
	responseResolver ResponseResolver[TId],
	responseSetter ResponseWriter,
) *DocumentRequestHandler[TDocument, TId] {
	this := new(DocumentRequestHandler[TDocument, TId])

	this.paramReader = paramReader
	this.responseResolver = responseResolver
	this.responseWriter = responseSetter

	return this
}

func (this DocumentRequestHandler[TDocument, TId]) Handle(ginContext *gin.Context) {
	id := this.paramReader.Read(ginContext)

	code, body := this.responseResolver.Resolve(id)

	this.responseWriter.Write(ginContext, code, body)
}
