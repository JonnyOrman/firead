package firead

import "cloud.google.com/go/firestore"

type SnapshotRetriever[TId Id] interface {
	Retrieve(collection *firestore.CollectionRef, id TId) *firestore.DocumentSnapshot
}
