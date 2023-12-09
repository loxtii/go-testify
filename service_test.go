package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	responseRecorder := getResponse("/cafe?count=2&city=moscow")

	// здесь нужно добавить необходимые проверки
	expectedCode := http.StatusOK
	expectedCount := 2
	body := strings.Split(responseRecorder.Body.String(), ",")

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code, "expected OK")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedBody := `wrong city value`
	responseRecorder := getResponse("/cafe?count=2&city=omsk")

	// здесь нужно добавить необходимые проверки
	expectedCode := http.StatusBadRequest
	body := responseRecorder.Body.String()

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code, "expected BadRequest")
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	responseRecorder := getResponse("/cafe?count=10&city=moscow")

	// здесь нужно добавить необходимые проверки
	expectedCode := http.StatusOK
	body := strings.Split(responseRecorder.Body.String(), ",")

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code, "expected OK")
	assert.Len(t, body, totalCount)
}

func TestMainHandlerWhenBadCount(t *testing.T) {
	responseRecorder := getResponse("/cafe?count=number&city=moscow")
	expectedCode := http.StatusBadRequest
	body := responseRecorder.Body.String()
	expectedBody := `wrong count value`

	assert.Equal(t, expectedCode, responseRecorder.Code, "expected BadRequest")
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenEmptyCount(t *testing.T) {
	responseRecorder := getResponse("/cafe?count=&city=moscow")
	expectedCode := http.StatusBadRequest
	body := responseRecorder.Body.String()
	expectedBody := `count missing`

	assert.Equal(t, expectedCode, responseRecorder.Code, "expected BadRequest")
	assert.Equal(t, expectedBody, body)
}
