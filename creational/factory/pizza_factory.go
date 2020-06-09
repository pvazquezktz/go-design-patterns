package factory

import (
	"errors"
	"fmt"
)

const (
	Veggie  = 1
	Special = 2
)

type IPizza interface {
	Make() string
}

func GetPizza(p int) (IPizza, error) {

	switch p {
	case Veggie:
		return new(VeggiePizza), nil
	case Special:
		return new(SpecialPizza), nil
	default:
		return nil, errors.New(fmt.Sprintf("Pizza type %v not recognized", p))
	}

}
