package firead

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type ParamReaderMock[TId Id] struct {
	mock.Mock
}

func (this ParamReaderMock[T]) Read(paramProvider ParamProvider) T {
	args := this.Called(paramProvider)
	return args.Get(0).(T)
}

type ResponseResolverMock[TId Id] struct {
	mock.Mock
}

func (this ResponseResolverMock[TId]) Resolve(id TId) (int, any) {
	args := this.Called(id)
	return args.Get(0).(int), args.Get(1).(any)
}

type ResponseWriterMock struct {
	mock.Mock
}

func (this ResponseWriterMock) Write(response Response, code int, body any) {
	this.Called(response, code, body)
}

func TestDocumentRequestHandlerHandle(t *testing.T) {
	id := "abc"

	code := 200

	document := make(map[string]interface{})

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	paramReader := new(ParamReaderMock[string])
	paramReader.On("Read", ginContext).Return(id)

	responseResolver := new(ResponseResolverMock[string])
	responseResolver.On("Resolve", id).Return(code, document)

	responseWriter := new(ResponseWriterMock)
	responseWriter.On("Write", ginContext, code, document).Return()

	sut := NewDocumentRequestHandler[map[string]interface{}, string](
		paramReader,
		responseResolver,
		responseWriter,
	)

	sut.Handle(ginContext)
}
