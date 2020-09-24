package facade

import "fmt"

// 提供食物的烧烤接口。
type Food interface {
	// 烧烤加料
	Barbecue()
}

// 通过接口暴露方法。单一职责,不必要的方法没有必要暴露。
type IAddOil interface {
	AddOil()
}

// 油
type Oil struct {
}

func (o *Oil) AddOil() {
	fmt.Println("add oil ...")
}

func (o *Oil) Test() {
	fmt.Println("add oil ...")
}

// 通过接口暴露方法。单一职责。不必要的方法没有必要暴露。
type IAddPepper interface {
	AddPepper()
}

// 花椒
type Pepper struct {
}

func (p *Pepper) AddPepper() {
	fmt.Println("add pepper ...")
}

// 通过接口暴露方法。单一职责。不必要的方法没有必要暴露。
type IAddHotPepper interface {
	AddHotPepper()
}

// 辣椒
type HotPepper struct {
}

func (p *HotPepper) AddHotPepper() {
	fmt.Println("add hotPepper ...")
}

type BarbecueFood struct {
}

//如果不太理解，你可以把他们想象成都是相互没有关系的模块，但是实现一个功能需要这几个模块的相互运算。
func (b *BarbecueFood) Barbecue() {

	// 这里实现烧烤的过程
	// 这里随便整修改都不影响调用的地方。

	//new(Oil).AddOil()
	//new(Pepper).AddPepper()
	//new(HotPepper).AddHotPepper()

	// 通过接口，只能调用暴露的方法。
	var iao IAddOil = new(Oil)
	iao.AddOil()
	// iao.Test()  报错，但确实有这个方法。
	var ip IAddPepper = new(Pepper)
	ip.AddPepper()
	var ihp IAddHotPepper = new(HotPepper)
	ihp.AddHotPepper()
}
