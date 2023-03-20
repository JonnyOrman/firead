package firead

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonExistingDocumentHandlerHandle(t *testing.T) {
	sut := NewNonExistingDocumentHandler()

	code, body := sut.Handle()

	assert.Equal(t, 404, code)
	assert.Equal(t, nil, body)
}
