package driven

type SendingSms interface {
	SendSms(phone string, content string) (bool, error)
}
