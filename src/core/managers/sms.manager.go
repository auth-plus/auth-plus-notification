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
	provider := e.chooseSmsProvider(phone, content)
	provider.SendSms(phone, content)
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *SmsManager) chooseSmsProvider(phone string, content string) d.SendingSms {
	if rand.Float64() < 0.5 {
		return e.onesignal
	}
	return e.sns
}
