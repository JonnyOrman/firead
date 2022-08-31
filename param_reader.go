package firead

type ParamReader[T any] interface {
	Read(paramProvider ParamProvider) T
}
