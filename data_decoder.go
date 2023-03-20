package firead

type DataDecoder[T any] interface {
	Decode(map[string]interface{}) T
}
