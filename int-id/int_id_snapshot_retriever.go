package intid

import (
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/firead"
)

type IntIdSnapshotRetriever struct {
	snapshotRetriever firead.SnapshotRetriever[string]
}

func NewIntIdSnapshotRetriever(
	snapshotRetriever firead.SnapshotRetriever[string],
) *IntIdSnapshotRetriever {
	this := new(IntIdSnapshotRetriever)

	this.snapshotRetriever = snapshotRetriever

	return this
}

func (this IntIdSnapshotRetriever) Retrieve(collection *firestore.CollectionRef, id int) *firestore.DocumentSnapshot {
	stringId := strconv.Itoa(id)

	return this.snapshotRetriever.Retrieve(collection, stringId)
}
