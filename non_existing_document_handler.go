package firead

type NonExistingDocumentHandler struct{}

func NewNonExistingDocumentHandler() *NonExistingDocumentHandler {
	this := new(NonExistingDocumentHandler)

	return this
}

func (this NonExistingDocumentHandler) Handle() (int, any) {
	return 404, make(map[string]interface{})
}
