package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVeggiePizza(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	pVeggie, err := pizzaCloner.GetClone(Veggie)
	assert.Nil(t, err, "err should be nil")
	msg := pVeggie.GetInfo()
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Eggplant, Mushroom - Shape: Circle - Dough: Gluten Free", msg, "msg is not the expected")

}

func TestGetSpecialPizza(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	sVeggie, err := pizzaCloner.GetClone(Special)
	assert.Nil(t, err, "err should be nil")
	msg := sVeggie.GetInfo()
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Ham, Pepper - Shape: Square - Dough: Regular", msg, "msg is not the expected")

}

func TestGetPizzaNonExistent(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	_, err := pizzaCloner.GetClone(50)
	assert.NotNil(t, err, "err should be nil")
	assert.Equal(t, "pizza type not recognized", err.Error(), "err is not the expected")

}
