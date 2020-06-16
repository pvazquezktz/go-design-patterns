package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPizzaFactory_Build(t *testing.T) {

	pizzaMaker, err := BuildFactory(PizzaFactoryType)
	assert.Nil(t, err, "err should be nil")

	vPizza, err2 := pizzaMaker.Build(Veggie)
	assert.Nil(t, err2, "err2 should be nil")

	msg := vPizza.MakeFood()
	assert.Equal(t, "Making a Veggie Pizza", msg, "msg is not the expected")

	sPizza, err3 := pizzaMaker.Build(Special)
	assert.Nil(t, err3, "err3 should be nil")

	msg2 := sPizza.MakeFood()
	assert.Equal(t, "Making a Special Pizza", msg2, "msg2 is not the expected")

}
