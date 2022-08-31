package firead

type StringIdReader struct{}

func NewStringIdReader() *StringIdReader {
	this := new(StringIdReader)

	return this
}

func (this StringIdReader) Read(paramProvider ParamProvider) string {
	id := paramProvider.Param("id")

	return id
}
