package firead

import "github.com/mitchellh/mapstructure"

type MapStructureDataDecoder[T any] struct{}

func NewMapStructureDataDecoder[T any]() *MapStructureDataDecoder[T] {
	this := new(MapStructureDataDecoder[T])

	return this
}

func (this MapStructureDataDecoder[T]) Decode(data map[string]interface{}) T {
	var result T
	mapstructure.Decode(data, &result)

	return result
}
