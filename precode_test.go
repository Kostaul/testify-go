package testing


import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"


	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)





func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req,err := http.NewRequest("GET","/count=5&city=moscow",nil)
	require.NoError(t, err, "Failed to create request")


    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
	response:=responseRecorder.Result()

	body, _ := io.ReadAll(response.Body)
	list := strings.Split(string(body), ",")

   assert.Len(t, list, totalCount, "Error - amount of cafes")
   require.Equal(t, 200, response.StatusCode, "Unexpected status code")
}



func TestMainHandlerWhenRequestIsValid(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
    require.NoError(t, err, "Failed to create request")

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    response := responseRecorder.Result()
    defer response.Body.Close()

    require.Equal(t, 200, response.StatusCode, "Unexpected status code")

    body, err := io.ReadAll(response.Body)
    assert.NoError(t, err)
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
    body, err := io.ReadAll(response.Body)
    assert.NoError(t, err)
    assert.Contains(t, string(body), "wrong city value")

}


