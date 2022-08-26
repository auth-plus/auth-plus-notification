# Providers

The provider is a struct that implements methods of 1 or multiples interfaces.

## Example

OneSignal can send an email, push notification, and SMS, this means that he could implement all three driven: SendingEmail, SendingPushNotification, and SendingSms

All implementation should be in here, and must not contain any business rule.
