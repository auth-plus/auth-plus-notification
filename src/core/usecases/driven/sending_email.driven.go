package driven

type SendingEmail interface {
	Send(email string, content string)
}
