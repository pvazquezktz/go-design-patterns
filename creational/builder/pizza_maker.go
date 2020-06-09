package builder

type IBuilder interface {
	SetIngredients() IBuilder
	SetDough() IBuilder
	SetShape() IBuilder
	Build() Pizza
}

// The director defines the order in which to execute the building steps
type PizzaMakerDirector struct {
	builder IBuilder
}

func (f *PizzaMakerDirector) Construct() Pizza {
	return f.builder.SetIngredients().SetDough().SetShape().Build()
}

// The builder provides the implementation for those steps.
func (f *PizzaMakerDirector) SetBuilder(b IBuilder) {
	f.builder = b
}

type Pizza struct {
	Ingredients []string
	Dough       string
	Shape       string
}
