package firead

import (
	"github.com/gin-gonic/gin"
)

type DocumentRequestHandler[TDocument any, TId any] struct {
	paramReader    ParamReader[TId]
	documentReader DocumentReader[TDocument, TId]
}

func NewDocumentRequestHandler[TDocument any, TId any](
	paramReader ParamReader[TId],
	documentReader DocumentReader[TDocument, TId],
) *DocumentRequestHandler[TDocument, TId] {
	this := new(DocumentRequestHandler[TDocument, TId])

	this.paramReader = paramReader
	this.documentReader = documentReader

	return this
}

func (this DocumentRequestHandler[TDocument, TId]) Handle(ginContext *gin.Context) {
	id := this.paramReader.Read(ginContext)

	this.documentReader.Read(id) //TODO - return document
}
