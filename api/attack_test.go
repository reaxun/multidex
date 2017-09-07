package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAttackFromAttacks(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/attacks/bubble")
	router.ServeHTTP(writer, req)
	expected := "{\"name\":\"bubble\",\"type\":\"water\",\"category\":\"special\",\"pp\":40,\"pow\":20,\"acc\":100}\n"
	assert.Equal(t, expected, writer.Body.String())
	assert.Equal(t, 200, writer.Code)
}

func TestGetAttackFromAttacksBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/attacks/bubble")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}

func TestGetAttackFromAttacksNotFound(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/attacks/badattack")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 404, writer.Code)
}

func TestGetAttacksByType(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/attacks/type/water")
	router.ServeHTTP(writer, req)
	expected := "{\"name\":\"bubble\",\"type\":\"water\",\"category\":\"special\",\"pp\":40,\"pow\":20,\"acc\":100}"
	assert.Contains(t, writer.Body.String(), expected)
	assert.Equal(t, 200, writer.Code)
}

func TestGetAttacksByTypeBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/attacks/type/water")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}

func TestGetAttacksByTypeNotFound(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/attacks/type/badtype")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 404, writer.Code)
}

func TestGetAttacks(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/attacks")
	router.ServeHTTP(writer, req)
	expected := "{\"name\":\"bubble\",\"type\":\"water\",\"category\":\"special\",\"pp\":40,\"pow\":20,\"acc\":100}"
	assert.Contains(t, writer.Body.String(), expected)
	assert.Equal(t, 200, writer.Code)
}

func TestGetAttacksBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/attacks")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}
