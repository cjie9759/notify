package wxroboot

import "testing"

var webhook = ""

func TestWXSend(t *testing.T) {
	n := NewNotify(Msgtype_text, webhook)
	n.Send("txtx test")

	n1 := NewNotify(Msgtype_markdown, webhook)
	n1.Send(`
	# md test
	**bold**
	[这是一个链接](http://work.weixin.qq.com/api/doc)`)

}
