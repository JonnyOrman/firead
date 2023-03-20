//go:build integration
// +build integration

package stringid

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"net/http"
)

var project = os.Getenv("PROJECT")
var firebaseEmulatorHost = os.Getenv("FIRESTORE_EMULATOR_HOST")
var appUrl = os.Getenv("APP_URL")

func TestMain(m *testing.M) {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, project)

	defer client.Close()

	collection := client.Collection("TestCollection")

	data := make(map[string]interface{})
	data["prop1"] = "def"
	data["prop2"] = 456

	collection.Doc("abc").Set(ctx, data)

	m.Run()

	documentsUrl := fmt.Sprintf("http://%s/emulator/v1/projects/%s/databases/(default)/documents", firebaseEmulatorHost, project)
	req, _ := http.NewRequest(http.MethodDelete, documentsUrl, nil)
	deleteClient := &http.Client{}
	deleteClient.Do(req)
}

func TestDocumentExists(t *testing.T) {
	url := fmt.Sprintf("%s/%s", appUrl, "abc")

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	expectedCode := 200
	expectedProp1 := "def"
	expectedProp2 := float64(456)

	assert.Equal(t, expectedCode, resp.StatusCode)
	assert.Equal(t, expectedProp1, m["prop1"])
	assert.Equal(t, expectedProp2, m["prop2"])
}

func TestDocumentDoesNotExist(t *testing.T) {
	url := fmt.Sprintf("%s/%s", appUrl, "ghi")

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	expectedCode := 404
	var expectedBody map[string]interface{}

	assert.Equal(t, expectedCode, resp.StatusCode)
	assert.Equal(t, expectedBody, m)
}
