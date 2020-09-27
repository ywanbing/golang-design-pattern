package singleton

import "sync"

type Cfg struct{}

// 饿汉单列 一开始就实例化
// 第一种方式，在能确定不会被并发影响的情况下使用
var CfgInstance1 = new(Cfg)
var CfgInstance2 *Cfg

var cfgInstance3 *Cfg
var cfgInstance4 *Cfg
var once sync.Once

// 饿汉单列 一开始就实例化
// 第二种方式 在能确定不会被并发影响的情况下使用
func init() {
	if CfgInstance2 == nil {
		CfgInstance2 = new(Cfg)
	}
}

// 懒汉单例 使用的时候实例化
// 第三种方式 在能确定不会被并发影响的情况下使用
func GetCfgInstance() *Cfg {
	if cfgInstance3 == nil {
		cfgInstance3 = new(Cfg)
	}
	return cfgInstance3
}

// 懒汉单例 使用的时候实例化
// 第四种方式,最佳。
func GetCfgInstance4() *Cfg {
	once.Do(func() {
		cfgInstance4 = new(Cfg)
	})
	return cfgInstance4
}
