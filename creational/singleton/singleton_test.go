package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {

	var pDelivery1, pDelivery2 ISingleton
	pDelivery1 = GetInstance()
	pDelivery2 = GetInstance()

	assert.Equal(t, pDelivery1, pDelivery2, "Instances are not the same")

	pDelivery1.AddOne()
	pDelivery2.AddOne()
	pizzasDelivered := pDelivery1.AddOne()

	assert.Equal(t, 3, pizzasDelivered, "pizzasDelivered is not the expected")

}
