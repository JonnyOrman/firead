package firead

type JsonResponseWriter struct{}

func NewJsonResponseWriter() *JsonResponseWriter {
	this := new(JsonResponseWriter)

	return this
}

func (this JsonResponseWriter) Write(response Response, code int, body any) {
	response.JSON(code, body)
}
