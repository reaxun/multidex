package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetAttackFromAttacks(t *testing.T) {
	writer := httptest.NewRecorder()
	router := mux.NewRouter()
	addRoutes(router)

	req, _ := getNewRequest("/attacks/bubble")
	router.ServeHTTP(writer, req)
	expected := "{\"name\":\"bubble\",\"type\":\"water\",\"category\":\"special\",\"pp\":40,\"pow\":20,\"acc\":100}\n"
	assert.Equal(t, expected, writer.Body.String())
}

// This needs to be fixed
//func TestGetAttackFromAttacksFail(t *testing.T) {
//	writer := httptest.NewRecorder()
//	router := mux.NewRouter()
//	addRoutes(router)
//
//	req, _ = getNewRequest("/attacks/badattack")
//	router.ServeHTTP(writer, req)
//	expected = "{}"
//	assert.Equal(t, expected, writer.Body.String())
//}
