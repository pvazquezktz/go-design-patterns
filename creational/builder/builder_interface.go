package builder

type IBuilder interface {
	SetIngredients() IBuilder
	SetDough() IBuilder
	SetShape() IBuilder
	Build() Pizza
}
