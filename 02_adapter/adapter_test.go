package adapter

import "testing"

func TestAdapter(t *testing.T) {

	// 正常播放
	var v Video = new(Mp4)
	v.Play()

	// 现在改用第三方开源库
	other := new(OtherMp4)
	v = new(Adapter).SetMp4(other)
	v.Play()

}
