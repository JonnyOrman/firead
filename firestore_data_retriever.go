package firead

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type FirestoreDataRetriever[TId Id] struct {
	configuration     fireworks.Configuration
	snapshotRetriever SnapshotRetriever[TId]
}

func NewFirestoreDataRetriever[TId Id](
	configuration fireworks.Configuration,
	snapshotRetriever SnapshotRetriever[TId]) *FirestoreDataRetriever[TId] {
	this := new(FirestoreDataRetriever[TId])

	this.configuration = configuration
	this.snapshotRetriever = snapshotRetriever

	return this
}

func (this FirestoreDataRetriever[TId]) Retrieve(id TId) map[string]interface{} {
	client, _ := firestore.NewClient(context.Background(), this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	snapshot := this.snapshotRetriever.Retrieve(collection, id)

	data := snapshot.Data()

	return data
}
