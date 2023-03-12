package notify

type Notify interface {
	Send(string) error
}

type NotifyGrop struct {
	notifys []Notify
}

func (n *NotifyGrop) Send(msg string) *NotifyGrop {
	for _, v := range n.notifys {
		go v.Send(msg)
	}
	return n
}

func NewNotifyGrop(notifys []Notify) *NotifyGrop {
	return &NotifyGrop{notifys}
}
