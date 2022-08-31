package firead

type DataRetriever[TId Id] interface {
	Retrieve(id TId) map[string]interface{}
}
