package firead

type ResponseResolver[TId Id] interface {
	Resolve(id TId) (int, any)
}
