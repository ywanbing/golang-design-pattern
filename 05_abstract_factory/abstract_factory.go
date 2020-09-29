package abstract_factory

import "fmt"

type Chair interface {
	SitOn()
}
type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("sit on modern chair ...")
}

type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("sit on victorian chair ...")
}

type Sofa interface {
	LyingDown()
}
type ModernSofa struct{}

func (m *ModernSofa) LyingDown() {
	fmt.Println("lying down modern sofa ...")
}

type VictorianSofa struct{}

func (v *VictorianSofa) LyingDown() {
	fmt.Println("lying down victorian sofa ...")
}

type Table interface {
	SomeThing()
}
type ModernTable struct{}

func (m *ModernTable) SomeThing() {
	fmt.Println("something modern table ...")
}

type VictorianTable struct{}

func (v *VictorianTable) SomeThing() {
	fmt.Println("something victorian table ...")
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateTable() Table
}

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) CreateChair() Chair {
	return new(ModernChair)
}

func (m *ModernFurnitureFactory) CreateSofa() Sofa {
	return new(ModernSofa)
}

func (m *ModernFurnitureFactory) CreateTable() Table {
	return new(ModernTable)
}

type VictorianFurnitureFactory struct{}

func (m *VictorianFurnitureFactory) CreateChair() Chair {
	return new(VictorianChair)
}

func (m *VictorianFurnitureFactory) CreateSofa() Sofa {
	return new(VictorianSofa)
}

func (m *VictorianFurnitureFactory) CreateTable() Table {
	return new(VictorianTable)
}
