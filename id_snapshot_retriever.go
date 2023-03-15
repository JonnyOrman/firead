package firead

import (
	"context"

	"cloud.google.com/go/firestore"
)

type IdSnapshotRetriever struct{}

func NewIdSnapshotRetriever() *IdSnapshotRetriever {
	this := new(IdSnapshotRetriever)

	return this
}

func (this IdSnapshotRetriever) Retrieve(collection *firestore.CollectionRef, id string) *firestore.DocumentSnapshot {
	ctx := context.Background()

	snapshot, _ := collection.Doc(id).Get(ctx)

	return snapshot
}
