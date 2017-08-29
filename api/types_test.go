package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeFromString(t *testing.T) {
	assert.Equal(t, Normal, TypeFromString("normal"))
	assert.Equal(t, None, TypeFromString("badtype"))
}

func TestGetMultiplier(t *testing.T) {
	assert.Equal(t, 1.0, GetMultiplier(Normal, Normal))
	assert.Equal(t, 0.0, GetMultiplier(Normal, Ghost))
	assert.Equal(t, 2.0, GetMultiplier(Water, Fire))
	assert.Equal(t, 0.5, GetMultiplier(Fire, Water))
}

func TestGetTypes(t *testing.T) {
	writer := httptest.NewRecorder()

	req, _ := getNewRequest("/types")
	router.ServeHTTP(writer, req)
	expected := "[\"normal\",\"fighting\",\"flying\",\"poison\",\"ground\",\"rock\",\"bug\",\"ghost\",\"steel\",\"fire\",\"water\",\"grass\",\"electric\",\"psychic\",\"ice\",\"dragon\",\"dark\",\"fairy\",\"none\"]\n"
	assert.Equal(t, expected, writer.Body.String())
	assert.Equal(t, 200, writer.Code)

}
