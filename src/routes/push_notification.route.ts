import { Router, Request, Response, NextFunction } from 'express'

import Core from '@core/layers'
import { PushNotificationDTO } from '@core/usecases/driver/send_push_notification.driver'

const pushNotificationRoute = Router()

pushNotificationRoute.post(
  '/',
  async (req: Request, res: Response, next: NextFunction) => {
    try {
      const data = new PushNotificationDTO(
        req.body.registrationToken,
        req.body.title,
        req.body.body,
        req.body.icon
      )
        .validate()
        .export()
      const resp = await Core.pushNotificationUsecase.sendToOne(data)
      res.status(200).send(resp)
    } catch (error) {
      next(error)
    }
  }
)

export default pushNotificationRoute
