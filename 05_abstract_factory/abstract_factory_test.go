package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {

	// 来个现代家具：
	var f FurnitureFactory = new(ModernFurnitureFactory)
	var chair Chair = f.CreateChair()
	chair.SitOn()
	var sofa Sofa = f.CreateSofa()
	sofa.LyingDown()
	var table Table = f.CreateTable()
	table.SomeThing()

	// 来个维多利亚 风格的家具
	f = new(VictorianFurnitureFactory)
	chair = f.CreateChair()
	chair.SitOn()
	sofa = f.CreateSofa()
	sofa.LyingDown()
	table = f.CreateTable()
	table.SomeThing()
}
