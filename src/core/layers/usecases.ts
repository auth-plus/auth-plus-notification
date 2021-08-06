import PushNotificationUsecase from '../usecases/push_notification.usecase'

import { firebaseServices } from './services'

export const pushNotificationUsecase = new PushNotificationUsecase(
  firebaseServices
)
