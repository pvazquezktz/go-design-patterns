package composition

import "fmt"

type Food struct {
	Name string
	Type string
}

func (f *Food) Show() string {
	return fmt.Sprintf("Food: %v (%v)", f.Name, f.Type)
}

type IngredientsList []string
