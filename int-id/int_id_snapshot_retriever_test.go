package intid

import (
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SnapshotRetrieverMock struct {
	mock.Mock
}

func (this SnapshotRetrieverMock) Retrieve(collection *firestore.CollectionRef, id string) *firestore.DocumentSnapshot {
	args := this.Called(collection, id)
	return args.Get(0).(*firestore.DocumentSnapshot)
}

func TestIntIdSnapshotRetrieverRetrieve(t *testing.T) {
	docRef := firestore.DocumentRef{}
	docRef.ID = "123"

	snapshot := firestore.DocumentSnapshot{}
	snapshot.Ref = &docRef

	collection := firestore.CollectionRef{}

	id := 123

	snapshotRetriever := new(SnapshotRetrieverMock)
	snapshotRetriever.On("Retrieve", &collection, "123").Return(&snapshot)

	sut := NewIntIdSnapshotRetriever(snapshotRetriever)

	result := sut.Retrieve(&collection, id)

	assert.Equal(t, snapshot.Ref, result.Ref)
}
