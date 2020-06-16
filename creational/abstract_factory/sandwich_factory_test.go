package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSandwichFactory_Build(t *testing.T) {

	sandwichMaker, err := BuildFactory(SandwichFactoryType)
	assert.Nil(t, err, "err should be nil")

	tSandwich, err2 := sandwichMaker.Build(Tuna)
	assert.Nil(t, err2, "err2 should be nil")

	msg := tSandwich.MakeFood()
	assert.Equal(t, "Making a Tuna Sandwich", msg, "msg is not the expected")

	cSandwich, err3 := sandwichMaker.Build(Cheese)
	assert.Nil(t, err3, "err3 should be nil")

	msg2 := cSandwich.MakeFood()
	assert.Equal(t, "Making a Cheese Sandwich", msg2, "msg2 is not the expected")
}
