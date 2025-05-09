package notify

import "bytes"

type Notify interface {
	Send(string) error
}

type NotifyGrop struct {
	notifys []Notify
	msgCh   chan string
	buf     bytes.Buffer
}

func NewNotifyGrop(notifys []Notify) *NotifyGrop {
	return &NotifyGrop{
		notifys: notifys,
		msgCh:   make(chan string, 100)}
}

func (n *NotifyGrop) Send(msg string) *NotifyGrop {
	for _, v := range n.notifys {
		go v.Send(msg)
	}
	return n
}
func (n *NotifyGrop) ChSend(msg string) *NotifyGrop {
	n.msgCh <- msg
	return n
}

func (n *NotifyGrop) ChSendDo() *NotifyGrop {
	defer n.buf.Reset()

	for {
		select {
		case v := <-n.msgCh:
			n.buf.WriteString(v)
			n.buf.WriteString("\r\n\r\n")
		default:
			n.Send(n.buf.String())
			return n
		}
	}

}
