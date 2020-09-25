package adapter

import "fmt"

type Video interface {
	Play()
}

type Mp4 struct {
}

// 本地 自己写的MP4 播放库。
func (m *Mp4) Play() {
	fmt.Println("play mp4 ...")
}

// 这种也可以
//type Adapter struct {
//	Other OtherMp4
//}
// 可以使用匿名嵌套，属于go语言的变种。
type Adapter struct {
	*OtherMp4
}

func (a *Adapter) Play() {
	// 这里的run 其实是 OtherMp4 的run方法。
	// 如果需要，可以重新为Adapter实现run 方法，会覆盖默认的run。
	a.Run()
}

// 类似这样的。
//func (a *Adapter) Run() {
// todo something
//a.OtherMp4.Run()
//}

func (a *Adapter) SetMp4(o *OtherMp4) Video {
	a.OtherMp4 = o
	return a
}

type OtherMp4 struct {
}

// 播放的方式和Mp4的播放不一样。引入第三方库。但是接口不一样。
func (m *OtherMp4) Run() {
	fmt.Println("run otherMp4 ...")
}
