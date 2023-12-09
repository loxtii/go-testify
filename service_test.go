package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedBody := `wrong city value`
	responseRecorder := getResponse("/cafe?count=2&city=omsk")

	// здесь нужно добавить необходимые проверки

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	responseRecorder := getResponse("/cafe?count=10&city=moscow")

	// здесь нужно добавить необходимые проверки

}
