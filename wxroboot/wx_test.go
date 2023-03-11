package wxroboot_test

import (
	"testing"

	"github.com/cjie9759/notify/wxroboot"
)

var webhook = ""

func TestWXSend(t *testing.T) {
	n := wxroboot.NewNotify(wxroboot.Msgtype_text, webhook)
	n.Send("txtx test")

	n1 := wxroboot.NewNotify(wxroboot.Msgtype_markdown, webhook)
	n1.Send(`
	# md test
	**bold**
	[这是一个链接](http://work.weixin.qq.com/api/doc)`)

}
