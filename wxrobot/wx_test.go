package wxrobot_test

import (
	"testing"

	"github.com/cjie9759/notify/wxrobot"
)

var webhook = ""

func TestWXSend(t *testing.T) {
	n := wxrobot.NewNotify(wxrobot.Msgtype_text, webhook, true)
	n.Send("txtx test")

	n1 := wxrobot.NewNotify(wxrobot.Msgtype_markdown, webhook, true)
	n1.Send(`
	# md test
	**bold**
	[这是一个链接](http://work.weixin.qq.com/api/doc)`)

}
