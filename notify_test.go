package notify_test

import (
	"testing"
	"time"

	"github.com/cjie9759/notify"
	"github.com/cjie9759/notify/mail"
	"github.com/cjie9759/notify/wxrobot"
)

var (
	MAIL_USER    = ""
	MAIL_PWD     = ""
	MAIL_FROM    = ""
	MAIL_TEST_TO = ""

	MAIL_SMTP_HOST = ""
	MAIL_SMTP_PORT = 0
	// MAIL_TO = []string{
	// 	"",
	// 	""}
	webhook     = ""
	CQ_GROUP_ID = 98
	CQ_URL      = ""
)

func TestSend(t *testing.T) {
	// os.Exit(1)
	notify.NewNotifyGrop([]notify.Notify{
		wxrobot.NewNotify(wxrobot.Msgtype_text, webhook, true),
		mail.NewMail(mail.Cfg{User: MAIL_USER,
			Pwd:      MAIL_PWD,
			From:     MAIL_FROM,
			To:       []string{MAIL_TEST_TO},
			Sub:      "ckie one mail test",
			STMPHOST: MAIL_SMTP_HOST,
		}),

		// cqrobot.NewNotify(CQ_GROUP_ID, CQ_URL),
	}).Send("group通知测试")
	time.Sleep(time.Minute)
}
func TestChSend(t *testing.T) {
	// os.Exit(1)
	notify.NewNotifyGrop([]notify.Notify{
		wxrobot.NewNotify(wxrobot.Msgtype_text, webhook, true),
		mail.NewMail(mail.Cfg{User: MAIL_USER,
			Pwd:      MAIL_PWD,
			From:     MAIL_FROM,
			To:       []string{MAIL_TEST_TO},
			Sub:      "ckie one mail test",
			STMPHOST: MAIL_SMTP_HOST,
		}),

		// cqrobot.NewNotify(CQ_GROUP_ID, CQ_URL),
	}).ChSend("msg1").ChSend("msg2").ChSendDo()

	// n.MsgCh

	time.Sleep(time.Minute)
}
