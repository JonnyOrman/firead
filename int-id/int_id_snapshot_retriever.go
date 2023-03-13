package intid

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type IntIdSnapshotRetriever struct {
	configuration fireworks.Configuration
}

func NewIntIdSnapshotRetriever(configuration fireworks.Configuration) *IntIdSnapshotRetriever {
	this := new(IntIdSnapshotRetriever)

	this.configuration = configuration

	return this
}

func (this IntIdSnapshotRetriever) Retrieve(collection *firestore.CollectionRef, id int) *firestore.DocumentSnapshot {
	ctx := context.Background()

	idString := strconv.Itoa(id)

	snapshot, _ := collection.Doc(idString).Get(ctx)

	return snapshot
}
