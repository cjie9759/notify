package notify

type Notify interface {
	Send(string) error
}
