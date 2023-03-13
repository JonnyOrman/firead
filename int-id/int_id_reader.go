package intid

import (
	"strconv"

	"github.com/jonnyorman/firead"
)

type IntIdReader struct{}

func NewIntIdReader() *IntIdReader {
	this := new(IntIdReader)

	return this
}

func (this IntIdReader) Read(paramProvider firead.ParamProvider) int {
	id := paramProvider.Param("id")

	idAsInt, _ := strconv.Atoi(id)

	return idAsInt
}
