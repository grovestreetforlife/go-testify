package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fullUrl(city string, count int) string {
	return fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
}

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	responseRecorder := getResponse(fullUrl("moscow", 2))

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedBody := `wrong city value` //в ТЗ на сайте просят "wrong count value", но это же к количеству ошибка
	responseRecorder := getResponse(fullUrl("omsk", 2))

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	responseRecorder := getResponse(fullUrl("moscow", 10))

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, totalCount)
}
