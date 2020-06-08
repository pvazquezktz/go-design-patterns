package builder

// Implements IBuilder interface
type VeggiePizzaBuilder struct {
	pizza Pizza
}

func (v *VeggiePizzaBuilder) SetIngredients() IBuilder {
	v.pizza.Ingredients = []string{"Tomato sauce", "Cheese", "Eggplant", "Mushroom"}
	return v
}

func (v *VeggiePizzaBuilder) SetDough() IBuilder {
	v.pizza.Dough = "Gluten Free"
	return v
}

func (v *VeggiePizzaBuilder) SetShape() IBuilder {
	v.pizza.Shape = "Circle"
	return v
}

func (v *VeggiePizzaBuilder) Build() Pizza {
	return v.pizza
}
