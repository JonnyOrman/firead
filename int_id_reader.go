package firead

import "strconv"

type IntIdReader struct{}

func NewIntIdReader() *IntIdReader {
	this := new(IntIdReader)

	return this
}

func (this IntIdReader) Read(paramProvider ParamProvider) int {
	id := paramProvider.Param("id")

	idAsInt, _ := strconv.Atoi(id)

	return idAsInt
}
