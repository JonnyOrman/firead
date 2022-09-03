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

type DocumentReaderMock[TDocument any, TId Id] struct {
	mock.Mock
}

func (this DocumentReaderMock[TDocument, TId]) Read(id TId) TDocument {
	args := this.Called(id)
	return args.Get(0).(TDocument)
}

type ResponseWriterMock struct {
	mock.Mock
}

func (this ResponseWriterMock) Write(response Response, code int, body any) {
	this.Called(response, code, body)
}

func TestDocumentRequestHandlerHandle(t *testing.T) {
	id := "abc"

	document := make(map[string]interface{})

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	paramReader := new(ParamReaderMock[string])
	paramReader.On("Read", ginContext).Return(id)

	documentReader := new(DocumentReaderMock[map[string]interface{}, string])
	documentReader.On("Read", id).Return(document)

	responseWriter := new(ResponseWriterMock)
	responseWriter.On("Write", ginContext, 200, document).Return()

	sut := NewDocumentRequestHandler[map[string]interface{}, string](
		paramReader,
		documentReader,
		responseWriter,
	)

	sut.Handle(ginContext)
}
