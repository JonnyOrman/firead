package firead

type ParamReader[T Id] interface {
	Read(paramProvider ParamProvider) T
}
