package builder

// Implements IBuilder interface
type SpecialPizzaBuilder struct {
	pizza Pizza
}

func (v *SpecialPizzaBuilder) SetIngredients() IBuilder {
	v.pizza.Ingredients = []string{"Tomato sauce", "Cheese", "Ham", "Pepper"}
	return v
}

func (v *SpecialPizzaBuilder) SetDough() IBuilder {
	v.pizza.Dough = "Regular"
	return v
}

func (v *SpecialPizzaBuilder) SetShape() IBuilder {
	v.pizza.Shape = "Square"
	return v
}

func (v *SpecialPizzaBuilder) Build() Pizza {
	return v.pizza
}
