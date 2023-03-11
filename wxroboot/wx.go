package wxroboot

import "github.com/imroc/req/v3"

type WXRNotify struct {
	msg      string
	config   Cfg
	msgtype  string
	msg_data map[string]any
}

func NewNotify(msgtype string) *WXRNotify {
	d := make(map[string]any, 3)
	return &WXRNotify{msgtype: msgtype, msg_data: d}
}

func (w *WXRNotify) Msg(msg string) *WXRNotify {
	w.msg = msg
	return w
}
func (w *WXRNotify) Send() error {
	w.msg_data["msgtype"] = w.msgtype
	// w.msg_data[w.msgtype] = make(map[string]string,1)
	w.msg_data[w.msgtype] = map[string]string{"content": w.msg}

	_, err := req.SetBodyJsonMarshal(w.msg_data).
		Post(w.config.webhook)
	return err
}

const (
	Msgtype_text     = "text"
	Msgtype_markdown = "markdown"
)

type Cfg struct {
	webhook string
}
