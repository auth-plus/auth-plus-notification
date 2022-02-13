package managers

import (
	d "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//Class for SmsManager
type SmsManager struct {
	sns       d.SendingSms
	onesignal d.SendingSms
}

func NewSmsManager(sns d.SendingSms, onesignal d.SendingSms) *SmsManager {
	return &SmsManager{sns: sns, onesignal: onesignal}
}

func (e *SmsManager) SendSms(phone string, content string) {
	choosedProvider := chooseSmsProvider(phone, content)
	switch choosedProvider {
	case "Sns":
		e.sns.SendSms(phone, content)
		break
	case "OneSignal":
		e.onesignal.SendSms(phone, content)
		break
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseSmsProvider(phone string, content string) EnumSmsProvider {
	if rand.Float64() < 0.5 {
		return "OneSignal"
	}
	return "Sns"
}
