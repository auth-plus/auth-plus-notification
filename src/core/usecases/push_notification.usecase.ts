import { SendingPushNotification } from './driven/sendingPushNotification.driven'
import {
  PushNotificationDTO,
  SendPushNotification,
  SendPushNotificationErrors,
} from './driver/send_push_notification.driver'

export default class PushNotificationUsecase implements SendPushNotification {
  constructor(private sendingPushNotification: SendingPushNotification) {}

  async sendToOne(data: PushNotificationDTO): Promise<void> {
    try {
      await this.sendingPushNotification.send(
        data.registrationToken,
        data.title,
        data.body,
        data.icon
      )
    } catch (error) {
      throw this.handleError(error)
    }
  }

  private handleError(error: Error) {
    if (error.message === 'PROVIDER_DEPENDENCY_ERROR') {
      return new SendPushNotificationErrors('DEPENDECY_ERROR')
    } else {
      return new SendPushNotificationErrors('ERROR_NOT_MAPPED')
    }
  }
}
