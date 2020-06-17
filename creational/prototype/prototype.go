package prototype

import "errors"

type IPizzaCloner interface {
	GetClone(c int) (IPizza, error)
}

const (
	Veggie  = 1
	Special = 2
)

type PizzaCloner struct{}

func (s *PizzaCloner) GetClone(c int) (IPizza, error) {

	switch c {
	case Veggie:
		newItem := *veggiePizza
		return &newItem, nil
	case Special:
		newItem := *specialPizza
		return &newItem, nil
	default:
		return nil, errors.New("pizza type not recognized")
	}

}
