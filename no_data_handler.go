package firead

type NoDataHandler interface {
	Handle() (int, any)
}
