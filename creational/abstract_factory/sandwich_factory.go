package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	Tuna   = 1
	Cheese = 2
)

type SandwichFactory struct{}

func (c *SandwichFactory) Build(v int) (IFood, error) {

	switch v {
	case Tuna:
		return new(TunaSandwich), nil
	case Cheese:
		return new(CheeseSandwich), nil
	default:
		return nil, errors.New(fmt.Sprintf("IFood of type %d not recognized", v))
	}

}
