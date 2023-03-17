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
var fireadStringIdAppUrl = os.Getenv("FIREAD_STRING_ID_APP_URL")

func TestMain(m *testing.M) {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, project)

	defer client.Close()

	collection := client.Collection("TestCollection")

	data := make(map[string]interface{})
	data["prop1"] = "val1"
	//data["prop2"] = 456

	collection.Doc("abc").Set(ctx, data)

	m.Run()

	documentsUrl := fmt.Sprintf("http://%s/emulator/v1/projects/%s/databases/(default)/documents", firebaseEmulatorHost, project)
	req, _ := http.NewRequest(http.MethodDelete, documentsUrl, nil)
	deleteClient := &http.Client{}
	deleteClient.Do(req)
}

func TestDocumentExists(t *testing.T) {
	url := fmt.Sprintf("http://firead-string-id-app:3001/%s", "abc")

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "val1", m["prop1"])
	//assert.Equal(t, 456, m["prop2"])
}

func TestDocumentDoesNotExist(t *testing.T) {
	url := fmt.Sprintf("%s/%s", fireadStringIdAppUrl, "456")
	fmt.Printf("URL: %s", url)

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	//assert.Equal(t, 404, resp.StatusCode)
	//assert.Equal(t, nil, m)
}
