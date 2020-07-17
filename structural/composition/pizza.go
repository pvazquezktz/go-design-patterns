package composition

type Pizza struct {
	Food                        // EMBEDDED composition. It's like inheritance (but it's not), fields and methods become part
	Ingredients IngredientsList // DIRECT composition, the field is a struct
}
