package simple_factory

import (
	"testing"
)

func TestFactory(t *testing.T) {

	food, err := BarbecueFactory("五花肉")
	if err != nil {
		panic(err)
	}
	food.Eaten()

	food, err = BarbecueFactory("羊肉串")
	if err != nil {
		panic(err)
	}
	food.Eaten()

	food, err = BarbecueFactory("")
	if err != nil {
		panic(err)
	}
	food.Eaten()

}
