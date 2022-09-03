package firead

type ParamProvider interface {
	Param(key string) string
}
