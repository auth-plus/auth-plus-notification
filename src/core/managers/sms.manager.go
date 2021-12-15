package managers

import (
	d "auth-plus-notification/core/usecases/driven"
)

//ENUM for Providers
type EnumSmsProvider string

const (
	SNS EnumSmsProvider = "Sns"
)

//Class for SmsManager
type SmsManager struct {
	sns d.SendingSms
}

func NewSmsManager(sns d.SendingSms) *SmsManager {
	return &SmsManager{sns: sns}
}

func (e *SmsManager) SendSms(phone string, content string) {
	choosedProvider := chooseSmsProvider(phone, content)
	switch choosedProvider {
	case "Sns":
		e.sns.SendSms(phone, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseSmsProvider(phone string, content string) EnumSmsProvider {
	return "Sns"
}
