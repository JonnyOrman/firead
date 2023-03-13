package firead

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type DocumentResponseResolver[TId Id] struct {
	configuration               *fireworks.Configuration
	snapshotRetriever           SnapshotRetriever[TId]
	existingDocumentDataHandler DataHandler[map[string]interface{}]
	nonExistingDocumentHandler  NoDataHandler
}

func NewDocumentResponseResolver[TId Id](
	configuration *fireworks.Configuration,
	snapshotRetriever SnapshotRetriever[TId],
	existingDocumentDataHandler DataHandler[map[string]interface{}],
	nonExistingDocumentHandler NoDataHandler,
) *DocumentResponseResolver[TId] {
	this := new(DocumentResponseResolver[TId])

	this.configuration = configuration
	this.snapshotRetriever = snapshotRetriever
	this.existingDocumentDataHandler = existingDocumentDataHandler
	this.nonExistingDocumentHandler = nonExistingDocumentHandler

	return this
}

func (this DocumentResponseResolver[TId]) Resolve(id TId) (int, any) {
	client, _ := firestore.NewClient(context.Background(), this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	snapshot := this.snapshotRetriever.Retrieve(collection, id)

	if snapshot != nil {
		return this.existingDocumentDataHandler.Handle(snapshot.Data())
	} else {
		return this.nonExistingDocumentHandler.Handle()
	}
}
