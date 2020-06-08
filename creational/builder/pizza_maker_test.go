package builder

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPizzaMakerDirector_Construct(t *testing.T) {

	pizzaMaker := PizzaMakerDirector{}

	veggiePizzaBuilder := new(VeggiePizzaBuilder)
	specialPizzaBuilder := new(SpecialPizzaBuilder)

	pizzaMaker.SetBuilder(veggiePizzaBuilder)
	vPizza := pizzaMaker.Construct()

	assert.Equal(t, strings.Join(vPizza.Ingredients, ","), "Tomato sauce,Cheese,Eggplant,Mushroom", "Ingredients are not the expected")
	assert.Equal(t, vPizza.Dough, "Gluten Free", "Dough are not the expected")
	assert.Equal(t, vPizza.Shape, "Circle", "Shape are not the expected")

	pizzaMaker.SetBuilder(specialPizzaBuilder)
	sPizza := pizzaMaker.Construct()

	assert.Equal(t, strings.Join(sPizza.Ingredients, ","), "Tomato sauce,Cheese,Ham,Pepper", "Ingredients are not the expected")
	assert.Equal(t, sPizza.Dough, "Regular", "Dough are not the expected")
	assert.Equal(t, sPizza.Shape, "Square", "Shape are not the expected")

}
