interface output {
  registrationToken: string
  title: string
  body: string
  icon?: string
}

export class PushNotificationDTO {
  constructor(
    public registrationToken: string,
    public title: string,
    public body: string,
    public icon?: string
  ) {}

  validate(): PushNotificationDTO {
    if (typeof this.registrationToken !== 'string') {
      throw new Error()
    }
    if (typeof this.title !== 'string') {
      throw new Error()
    }
    if (typeof this.body !== 'string') {
      throw new Error()
    }
    if (typeof this.icon !== 'string') {
      throw new Error()
    }
    return this
  }

  export(): PushNotificationDTO {
    return {
      registrationToken: this.registrationToken,
      title: this.title,
      body: this.body,
      icon: this.icon,
    } as PushNotificationDTO
  }
}

export interface SendPushNotification {
  sendToOne: (data: PushNotificationDTO) => Promise<void>
}

export type SendPushNotificationErrorsTypes =
  | 'DEPENDECY_ERROR'
  | 'ERROR_NOT_MAPPED'

export class SendPushNotificationErrors extends Error {
  constructor(message: SendPushNotificationErrorsTypes) {
    super(message)
    this.name = 'LoginUser'
  }
}
