package firead

type ResponseWriter interface {
	Write(response Response, code int, body any)
}
