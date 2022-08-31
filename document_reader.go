package firead

type DocumentReader[TDocument any, TId any] interface {
	Read(id TId) TDocument
}
