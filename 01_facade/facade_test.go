package facade

import "testing"

func TestFood(t *testing.T) {
	// 通过接口调用
	// 简单的一步封装搞定一系列的小功能。
	var f Food = new(BarbecueFood)
	f.Barbecue()
}
