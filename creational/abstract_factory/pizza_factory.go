package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	Veggie  = 1
	Special = 2
)

type PizzaFactory struct{}

func (c *PizzaFactory) Build(v int) (IFood, error) {

	switch v {
	case Veggie:
		return new(VeggiePizza), nil
	case Special:
		return new(SpecialPizza), nil
	default:
		return nil, errors.New(fmt.Sprintf("IFood of type %d not recognized", v))
	}

}
