package prototype

import (
	"fmt"
	"strings"
)

type IPizza interface {
	GetInfo() string
}

type Pizza struct {
	Ingredients []string
	Dough       string
	Shape       string
}

func (p *Pizza) GetInfo() string {
	return fmt.Sprintf("Contains: %v - Shape: %v - Dough: %v", strings.Join(p.Ingredients, ", "), p.Shape, p.Dough)
}

var veggiePizza = &Pizza{
	Ingredients: []string{"Tomato sauce", "Cheese", "Eggplant", "Mushroom"},
	Dough:       "Gluten Free",
	Shape:       "Circle",
}

var specialPizza = &Pizza{
	Ingredients: []string{"Tomato sauce", "Cheese", "Ham", "Pepper"},
	Dough:       "Regular",
	Shape:       "Square",
}
