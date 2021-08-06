export interface SendingPushNotification {
  send: (
    registrationToken: string,
    title: string,
    body: string,
    icon?: string
  ) => Promise<void>
}

export type SendingPushNotificationErrorsTypes = 'PROVIDER_DEPENDENCY_ERROR'

export class SendingPushNotificationErrors extends Error {
  constructor(message: SendingPushNotificationErrorsTypes) {
    super(message)
  }
}
