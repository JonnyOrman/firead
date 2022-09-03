package firead

type Response interface {
	JSON(code int, obj any)
}
