package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {

	// 来个A
	var fa Factory = new(FactoryA)
	var a Product = fa.NewProduct()
	a.Show()

	// 来个B
	var fb Factory = new(FactoryB)
	var b Product = fb.NewProduct()
	b.Show()
}
