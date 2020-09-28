package factory_method

import "fmt"

// 商品接口
type Product interface {
	Show()
}

type A struct{}

func (a *A) Show() {
	fmt.Println("生产 A 产品。。。。")
}

type B struct{}

func (b *B) Show() {
	fmt.Println("生产 B 产品。。。。")
}

// 工厂接口
type Factory interface {
	NewProduct() Product
}

type FactoryA struct{}

func (f *FactoryA) NewProduct() Product {
	return new(A)
}

type FactoryB struct{}

func (f *FactoryB) NewProduct() Product {
	return new(B)
}
