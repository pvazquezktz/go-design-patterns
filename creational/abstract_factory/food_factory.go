package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	PizzaFactoryType    = 1
	SandwichFactoryType = 2
)

type IFood interface {
	MakeFood() string
}

type IFoodFactory interface {
	Build(v int) (IFood, error)
}

func BuildFactory(f int) (IFoodFactory, error) {
	switch f {
	case PizzaFactoryType:
		return new(PizzaFactory), nil
	case SandwichFactoryType:
		return new(SandwichFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized", f))
	}
}
