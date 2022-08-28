package driven

type SendingEmail interface {
	SendEmail(email string, content string) (bool, error)
}
