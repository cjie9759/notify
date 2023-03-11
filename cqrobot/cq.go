package cqrobot

import (
	"strconv"

	"github.com/imroc/req/v3"
)

type CQRNotify struct {
	// msg      string
	url      string
	group_id int
	msg_data map[string]string
}

func NewNotify(group_id int, url string) *CQRNotify {
	// d := make(map[string]string, 3)
	return &CQRNotify{
		url,
		group_id,
		map[string]string{
			"group_id":    strconv.Itoa(group_id),
			"message":     "defalt msg",
			"auto_escape": "false",
		},
	}
}

func (w *CQRNotify) Send(msg string) error {

	w.msg_data["message"] = msg

	_, err := req.SetBodyJsonMarshal(w.msg_data).
		Post(w.url)
	return err
}

const (
	Msgtype_text     = "text"
	Msgtype_markdown = "markdown"
)
