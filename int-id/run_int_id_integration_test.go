package intid

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestMain(m *testing.M) {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "firebase-emulator:8080")
	os.Setenv("PROJECT", "firead-int-id-integration-tests")
	os.Setenv("PROJECT_ID", "firead-int-id-integration-tests")

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "firead-int-id-integration-tests")
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
	}

	defer client.Close()

	collection := client.Collection("TestCollection")

	data := make(map[string]interface{})
	data["prop1"] = "val1"
	data["prop2"] = 456

	collection.Doc("123").Set(ctx, data)

	m.Run()

	req, _ := http.NewRequest(http.MethodDelete, "http://firebase-emulator:8080/emulator/v1/projects/firead-int-id-integration-tests/databases/(default)/documents", nil)
	deleteClient := &http.Client{}
	deleteClient.Do(req)
}

func TestDocumentExists(t *testing.T) {
	url := fmt.Sprintf("http://firead-int-id-app:3001/%s", "123")
	fmt.Printf("URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
	}

	log.Printf("RESP BODY: %s", body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "val1", m["prop1"])
	//assert.Equal(t, 456, m["prop2"])
}

func TestDocumentDoesNotExist(t *testing.T) {
	url := fmt.Sprintf("http://firead-int-id-app:3001/%s", "456")
	fmt.Printf("URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
	}

	log.Printf("RESP BODY: %s", body)

	var m map[string]interface{}
	json.Unmarshal(body, &m)

	//assert.Equal(t, 404, resp.StatusCode)
	//assert.Equal(t, nil, m)
}
