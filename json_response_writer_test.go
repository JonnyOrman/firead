package firead

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type ResponseMock struct {
	mock.Mock
}

func (this ResponseMock) JSON(code int, body any) {
	this.Called(code, body)
}

func TestJsonResponseWriterWrite(t *testing.T) {
	code := 200

	body := TestDocument{}
	body.Prop1 = "abc"
	body.Prop2 = 123

	response := new(ResponseMock)
	response.On("JSON", code, body).Return()

	sut := NewJsonResponseWriter()

	sut.Write(response, code, body)
}
