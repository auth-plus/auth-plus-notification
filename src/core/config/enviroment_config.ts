import * as dotenv from 'dotenv'

dotenv.config()

export default {
  app: {
    name: process.env.APP_NAME ?? 'auth-plus-notification',
    port: parseInt(process.env.APP_PORT ?? '5001'),
    enviroment: process.env.NODE_ENV ?? 'development',
  },
  firebase: {
    databaseURL: process.env.FIREBASE_DATABASE_URL,
  },
}
