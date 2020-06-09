package builder

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPizzaMakerWithoutDirector(t *testing.T) {

	pizzaMaker := PizzaMakerBuilder{}
	vPizza := pizzaMaker.SetIngredients([]string{"Tomato sauce", "Cheese", "Eggplant", "Mushroom"}).
		SetDough("Gluten Free").SetShape("Circle").Build()

	assert.Equal(t, strings.Join(vPizza.Ingredients, ","), "Tomato sauce,Cheese,Eggplant,Mushroom", "Ingredients are not the expected")
	assert.Equal(t, vPizza.Dough, "Gluten Free", "Dough are not the expected")
	assert.Equal(t, vPizza.Shape, "Circle", "Shape are not the expected")

	pizzaMaker2 := PizzaMakerBuilder{}
	sPizza := pizzaMaker2.SetIngredients([]string{"Tomato sauce", "Cheese", "Ham", "Pepper"}).
		SetDough("Regular").SetShape("Square").Build()

	assert.Equal(t, strings.Join(sPizza.Ingredients, ","), "Tomato sauce,Cheese,Ham,Pepper", "Ingredients are not the expected")
	assert.Equal(t, sPizza.Dough, "Regular", "Dough are not the expected")
	assert.Equal(t, sPizza.Shape, "Square", "Shape are not the expected")

}
