import { messaging } from 'firebase-admin'

import { msg } from '../config/firebase'

import logger from '@core/config/logger'
import {
  SendingPushNotification,
  SendingPushNotificationErrors,
} from '@core/usecases/driven/sendingPushNotification.driven'

export class FirebaseServices implements SendingPushNotification {
  async send(
    registrationToken: string,
    title: string,
    body: string,
    icon?: string
  ): Promise<void> {
    try {
      const payload: messaging.MessagingPayload = {
        notification: {
          title,
          body,
          icon,
        },
      }
      msg.sendToDevice(registrationToken, payload)
    } catch (error) {
      logger.error(error)
      throw new SendingPushNotificationErrors('PROVIDER_DEPENDENCY_ERROR')
    }
  }
}
