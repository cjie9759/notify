package mail_test

import (
	"testing"

	"github.com/cjie9759/notify/mail"
)

var (
	MAIL_USER    = ""
	MAIL_PWD     = ""
	MAIL_FROM    = ""
	MAIL_TEST_TO = ""
	MAIL_TO      = []string{}
)

func TestXxx(t *testing.T) {
	// mail1 := mail.NewMail(MAIL_USER, MAIL_PWD, MAIL_FROM, []string{MAIL_TEST_TO}, "ckie onen mail test")
	mail1 := mail.NewMail(mail.Cfg{MAIL_USER, MAIL_PWD, MAIL_FROM, []string{MAIL_TEST_TO}, "ckie onen mail test"})
	mail1.Send("ckie onen mail test mail test")
}
