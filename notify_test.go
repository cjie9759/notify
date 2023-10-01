package notify_test

import (
	"testing"
	"time"

	"github.com/cjie9759/notify"
	"github.com/cjie9759/notify/cqrobot"
	"github.com/cjie9759/notify/mail"
	"github.com/cjie9759/notify/wxrobot"
)

var (
	MAIL_USER = ""
	MAIL_PWD  = ""
	MAIL_FROM = ""

	MAIL_TEST_TO = ""

	MAIL_TO = []string{
		"",
		""}
	webhook     = ""
	CQ_GROUP_ID = 98
	CQ_URL      = ""
)

func TestXxx(t *testing.T) {
	// os.Exit(1)
	notify.NewNotifyGrop([]notify.Notify{
		wxrobot.NewNotify(wxrobot.Msgtype_text, webhook),
		mail.NewMail(mail.Cfg{User: MAIL_USER, Pwd: MAIL_PWD, From: MAIL_FROM, To: []string{MAIL_TEST_TO}, Sub: "ckie onen mail test"}),
		cqrobot.NewNotify(CQ_GROUP_ID, CQ_URL),
	}).Send("group通知测试")
	time.Sleep(time.Minute)
}
