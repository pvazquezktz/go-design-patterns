package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVeggiePizza(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	pVeggie, err := pizzaCloner.GetClone(Veggie)
	pVeggie2, err2 := pizzaCloner.GetClone(Veggie)

	assert.Nil(t, err, "err should be nil")
	assert.Nil(t, err2, "err2 should be nil")

	// Not the same instance
	assert.True(t, pVeggie != pVeggie2, "pVeggie should not be equal to pVeggie2")

	msg := pVeggie.GetInfo()
	msg2 := pVeggie2.GetInfo()
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Eggplant, Mushroom - Shape: Circle - Dough: Gluten Free", msg, "msg is not the expected")
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Eggplant, Mushroom - Shape: Circle - Dough: Gluten Free", msg2, "msg2 is not the expected")

}

func TestGetSpecialPizza(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	pSpecial, err := pizzaCloner.GetClone(Special)
	pSpecial2, err2 := pizzaCloner.GetClone(Special)

	assert.Nil(t, err, "err should be nil")
	assert.Nil(t, err2, "err2 should be nil")

	// Not the same instance
	assert.True(t, pSpecial != pSpecial2, "pSpecial should not be equal to pSpecial2")

	msg := pSpecial.GetInfo()
	msg2 := pSpecial2.GetInfo()
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Ham, Pepper - Shape: Square - Dough: Regular", msg, "msg is not the expected")
	assert.Equal(t, "Contains: Tomato sauce, Cheese, Ham, Pepper - Shape: Square - Dough: Regular", msg2, "msg2 is not the expected")

}

func TestGetPizzaNonExistent(t *testing.T) {

	pizzaCloner := PizzaCloner{}

	_, err := pizzaCloner.GetClone(50)
	assert.NotNil(t, err, "err should be nil")
	assert.Equal(t, "pizza type not recognized", err.Error(), "err is not the expected")

}
