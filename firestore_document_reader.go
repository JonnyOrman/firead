package firead

import (
	"github.com/mitchellh/mapstructure"
)

type FirestoreDocumentReader[TDocument any, TId Id] struct {
	dataRetriever DataRetriever[TId]
}

func NewFirestoreDocumentReader[TDocument any, TId Id](
	dataRetriever DataRetriever[TId]) *FirestoreDocumentReader[TDocument, TId] {
	this := new(FirestoreDocumentReader[TDocument, TId])

	this.dataRetriever = dataRetriever

	return this
}

func (this FirestoreDocumentReader[TDocument, TId]) Read(id TId) TDocument {
	data := this.dataRetriever.Retrieve(id)

	var result TDocument
	mapstructure.Decode(data, &result)

	return result
}
