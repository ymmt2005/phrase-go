/*
Phrase Strings API Reference

Testing JobsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package phrase

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phrase/phrase-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_phrase_JobsApiService(t *testing.T) {
	t.Run("Test JobsApiService JobCreate", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// assert request body
			var requestBody map[string]interface{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&requestBody); err != nil {
				t.Fatal(err)
			}

			// Access the values from the JSON payload
			assert.Equal(t, requestBody["name"], "createjobactionb1")
			assert.Equal(t, requestBody["due_date"], nil)

			// Send the mock response
			response := `{"id": "1", "name": "createjobactionb1", "code": "123abc", "state": "draft", "user_id": "1" }`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(response))
		}))

		defer server.Close()

		configuration := phrase.NewConfiguration()
		configuration.BasePath = server.URL
		apiClient := phrase.NewAPIClient(configuration)

		translationKeyIds := []string{"1", "2", "3"}

		localVarOptionals := phrase.JobCreateOpts{}
		jobCreateParams := phrase.JobCreateParameters{Name: "createjobactionb1", TranslationKeyIds: translationKeyIds}
		resp, httpRes, err := apiClient.JobsApi.JobCreate(context.Background(), "project_id", jobCreateParams, &localVarOptionals)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}
