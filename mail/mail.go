package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

type Mail struct {
	// msg    string
	cfg Cfg
}

type Cfg struct {
	User string
	Pwd  string
	From string
	To   []string
	Sub  string

	STMPHOST string
	STMPPORT int
}

var defaultCfg = Cfg{
	STMPHOST: "smtp.exmail.qq.com",
	STMPPORT: 465,
}

func NewMail(config Cfg) *Mail {

	m := &Mail{cfg: config}

	if m.cfg.STMPHOST == "" {
		m.cfg.STMPHOST = defaultCfg.STMPHOST
	}
	if m.cfg.STMPPORT == 0 {
		m.cfg.STMPPORT = defaultCfg.STMPPORT
	}

	return m
}

// func (M *Mail) Send(Body string) error {
// 	M.msg = Body
// 	return M.send()
// }

func (M *Mail) Send(Body string) error {

	// 换行替换
	// body := M.msg
	// body := strings.ReplaceAll(M.Body, "\n", "<\br>")

	msgBuffer := bytes.Buffer{}
	// bytes.NewBuffer([]byte(fmt.Sprintf()))
	msgBuffer.WriteString(fmt.Sprintf("From:%s<%s>\r\n", M.cfg.From, M.cfg.User))
	msgBuffer.WriteString(fmt.Sprintf("To:%s\r\n", strings.Join(M.cfg.To, ",")))
	msgBuffer.WriteString(fmt.Sprintf("Subject:%s\r\n", M.cfg.Sub))
	msgBuffer.WriteString("Content-Type:text/plain;chartset=UTF-8\r\n\r\n")
	msgBuffer.WriteString(Body)

	auth := smtp.PlainAuth(
		"",
		M.cfg.User,
		M.cfg.Pwd,
		M.cfg.STMPHOST,
	)

	err := M.sendMailUsingTLS(
		fmt.Sprintf("%s:%d", M.cfg.STMPHOST, M.cfg.STMPPORT),
		auth,
		M.cfg.User,
		M.cfg.To,
		msgBuffer.Bytes(),
	)
	// msgBuffer.Reset()
	return err
	// if err != nil {
	// 	panic(err)
	// }

}

// return a smtp client
func (M *Mail) dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		// log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// 参考net/smtp的func (M *Mail) SendMail()
// 使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
// len(to)>1时,to[1]开始提示是密送
func (M *Mail) sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := M.dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
