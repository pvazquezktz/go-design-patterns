package singleton

import (
	"sync"
)

var instance *PizzaDelivery
var mux sync.Mutex

// PizzaDelivery implements ISingleton interface
type PizzaDelivery struct {
	Delivered int
}

func (s *PizzaDelivery) AddOne() int {
	s.Delivered++
	return s.Delivered
}

func GetInstance() ISingleton {
	if instance == nil {
		// This validations guarantees the function is thread safe
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = new(PizzaDelivery)
		}
	}
	return instance
}
