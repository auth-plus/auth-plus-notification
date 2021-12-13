package driven

type SendingSms interface {
	Send(phone string, content string)
}
