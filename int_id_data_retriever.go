package firead

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type IntIdDataRetriever struct {
	configuration fireworks.Configuration
}

func NewIntIdDataRetriever(configuration fireworks.Configuration) *IntIdDataRetriever {
	this := new(IntIdDataRetriever)

	this.configuration = configuration

	return this
}

func (this IntIdDataRetriever) Retrieve(id int) map[string]interface{} {
	client, _ := firestore.NewClient(context.Background(), this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	snapshot, _ := collection.Doc(string(id)).Get(context.Background())

	data := snapshot.Data()

	return data
}
