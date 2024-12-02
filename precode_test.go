package mainn

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req, err := http.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
    require.NoError(t, err, "Failed to create request")

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    response := responseRecorder.Result()
    defer response.Body.Close()

    require.Equal(t, http.StatusOK, response.StatusCode, "Unexpected status code")

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    assert.Len(t, list, totalCount, "Error - amount of cafes")
}

func TestMainHandlerWhenRequestIsValid(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
    require.NoError(t, err, "Failed to create request")

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    response := responseRecorder.Result()
    defer response.Body.Close()

    require.Equal(t, http.StatusOK, response.StatusCode, "Unexpected status code")

    body := responseRecorder.Body.String()
    assert.NotEmpty(t, body, "Response body is empty")
}

func TestMainHandlerWhenCityIsUnsupported(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=unknowncity", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    response := responseRecorder.Result()
    defer response.Body.Close()

    require.Equal(t, http.StatusBadRequest, response.StatusCode, "Unexpected status code")

    body := responseRecorder.Body.String()
    assert.Contains(t, body, "wrong city value")
}
