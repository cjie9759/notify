package cqrobot_test

import (
	"testing"

	"github.com/cjie9759/notify/cqrobot"
)

var (
	CQ_GROUP_ID = 12323
	CQ_URL      = ""
)

func TestCqSend(t *testing.T) {
	n := cqrobot.NewNotify(CQ_GROUP_ID, CQ_URL)
	n.Send("txtx test")

	n.Send(`
	# md test
	**bold**
	[这是一个链接](http://work.weixin.qq.com/api/doc)`)

}
