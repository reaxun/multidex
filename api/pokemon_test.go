package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonFromPokedex(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/pokemon/pikachu")
	router.ServeHTTP(writer, req)
	expected := "{\"number\":25,\"name\":\"pikachu\",\"basestats\":{\"hp\":35,\"attack\":55,\"defense\":40,\"sattack\":50,\"sdefense\":50,\"speed\":90},\"typea\":\"electric\",\"typeb\":\"none\"}\n"
	assert.Equal(t, expected, writer.Body.String())
	assert.Equal(t, 200, writer.Code)
}

func TestGetPokemonFromPokedexBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/pokemon/pikachu")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}

func TestGetPokemonFromPokedexNotFound(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/pokemon/badpokemon")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 404, writer.Code)
}

func TestGetPokemonByType(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/pokemon/type/electric")
	router.ServeHTTP(writer, req)
	expected := "{\"number\":25,\"name\":\"pikachu\",\"basestats\":{\"hp\":35,\"attack\":55,\"defense\":40,\"sattack\":50,\"sdefense\":50,\"speed\":90},\"typea\":\"electric\",\"typeb\":\"none\"}"
	assert.Contains(t, writer.Body.String(), expected)
	assert.Equal(t, 200, writer.Code)
}

func TestGetPokemonByTypeBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/pokemon/type/electric")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}

func TestGetPokemonByTypeNotFound(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/pokemon/type/badtype")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 404, writer.Code)
}

func TestGetPokemon(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/pokemon")
	router.ServeHTTP(writer, req)
	expected := "{\"number\":25,\"name\":\"pikachu\",\"basestats\":{\"hp\":35,\"attack\":55,\"defense\":40,\"sattack\":50,\"sdefense\":50,\"speed\":90},\"typea\":\"electric\",\"typeb\":\"none\"}"
	assert.Contains(t, writer.Body.String(), expected)
	assert.Equal(t, 200, writer.Code)
}

func TestGetPokemonBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	openFakeDB()

	req, _ := getNewRequest("/pokemon")
	router.ServeHTTP(writer, req)
	assert.Equal(t, 400, writer.Code)
	closeFakeDB()
}
