package simple_factory

import "fmt"

type Barbecue interface {
	// 被吃
	Eaten()
}

type 羊肉串 struct {
}

func (y 羊肉串) Eaten() {
	fmt.Println("我吃过羊肉串了")
}

type 五花肉 struct {
}

func (w 五花肉) Eaten() {
	fmt.Println("我吃过五花肉了")
}

// 创建一个工厂。不一定必须要新建一个结构体来管理，也可以通过函数来管理工厂
func BarbecueFactory(name string) (Barbecue, error) {
	switch name {
	case "羊肉串":
		y := new(羊肉串)
		// todo 加工一下
		return y, nil
	case "五花肉":
		w := new(五花肉)
		// todo 加工一下
		return w, nil
	default:
		return nil, fmt.Errorf("没有这东西")
	}
}
