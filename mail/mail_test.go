package mail_test

import (
	"fmt"
	"testing"

	"github.com/cjie9759/notify/mail"
)

var (
	MAIL_USER    = ""
	MAIL_PWD     = ""
	MAIL_FROM    = ""
	MAIL_TEST_TO = ""
	MAIL_TO      = []string{}

	MAIL_SMTP_HOST = ""
	MAIL_SMTP_PORT = 0
)

func TestXxx(t *testing.T) {
	// mailTest := mail.NewMail(MAIL_USER, MAIL_PWD, MAIL_FROM, []string{MAIL_TEST_TO}, "ckie one mail test")
	mailTest := mail.NewMail(
		mail.Cfg{
			MAIL_USER,
			MAIL_PWD,
			MAIL_FROM,
			[]string{MAIL_TEST_TO},
			"ckie one mail test", MAIL_SMTP_HOST, 0})

	err := mailTest.Send("ckie one mail test mail test")
	if err != nil {
		fmt.Println(err)
	}
}
