package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVeggiePizza(t *testing.T) {

	vPizza, err := GetPizza(Veggie)

	assert.Nil(t, err, "err should be nil")
	msg := vPizza.Make()
	assert.Equal(t, "Making a Veggie Pizza", msg, "msg is not the expected")

}

func TestGetSpecialPizza(t *testing.T) {

	vPizza, err := GetPizza(Special)

	assert.Nil(t, err, "err should be nil")
	msg := vPizza.Make()
	assert.Equal(t, "Making a Special Pizza", msg, "msg is not the expected")

}

func TestGetPizzaNonExistent(t *testing.T) {

	_, err := GetPizza(50)

	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "Pizza type 50 not recognized", err.Error(), "err is not the expected")

}
