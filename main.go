package notify

type Notify interface {
	Msg(string) *Notify
	Send() error
}
