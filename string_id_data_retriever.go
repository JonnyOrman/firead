package firead

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type StringIdDataRetriever struct {
	configuration fireworks.Configuration
}

func NewStringIdDataRetriever(configuration fireworks.Configuration) *StringIdDataRetriever {
	this := new(StringIdDataRetriever)

	this.configuration = configuration

	return this
}

func (this StringIdDataRetriever) Retrieve(id string) map[string]interface{} {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	snapshot, _ := collection.Doc(id).Get(ctx)

	data := snapshot.Data()

	return data
}
