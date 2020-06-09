package builder

type PizzaMakerBuilder struct {
	pizza Pizza
}

func (p *PizzaMakerBuilder) SetIngredients(ingredients []string) *PizzaMakerBuilder {
	p.pizza.Ingredients = append(p.pizza.Ingredients, ingredients...)
	return p
}

func (p *PizzaMakerBuilder) SetDough(dough string) *PizzaMakerBuilder {
	p.pizza.Dough = dough
	return p
}

func (p *PizzaMakerBuilder) SetShape(shape string) *PizzaMakerBuilder {
	p.pizza.Shape = shape
	return p
}

func (p *PizzaMakerBuilder) Build() Pizza {
	return p.pizza
}
