package driven

type PushNotificatioManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingPushNotification, error)
}
