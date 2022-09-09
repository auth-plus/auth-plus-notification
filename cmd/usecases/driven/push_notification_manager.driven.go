package driven

// PushNotificatioManager is a interface that must abstract whats input should be used for choosing a provider
type PushNotificatioManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingPushNotification, error)
}
