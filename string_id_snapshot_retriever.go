package firead

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type StringIdSnapshotRetriever struct {
	configuration fireworks.Configuration
}

func NewStringIdSnapshotRetriever(configuration fireworks.Configuration) *StringIdSnapshotRetriever {
	this := new(StringIdSnapshotRetriever)

	this.configuration = configuration

	return this
}

func (this StringIdSnapshotRetriever) Retrieve(collection *firestore.CollectionRef, id string) *firestore.DocumentSnapshot {
	ctx := context.Background()

	snapshot, _ := collection.Doc(id).Get(ctx)

	return snapshot
}
