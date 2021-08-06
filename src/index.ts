import cors from 'cors'
import express, {
  Request,
  Response,
  NextFunction,
  urlencoded,
  json,
} from 'express'
import helmet from 'helmet'

import pushNotificationRoute from './routes/push_notification.route'

import config from '@core/config/enviroment_config'
import logger from '@core/config/logger'

const app = express()

app.use(helmet())
app.use(
  cors({
    origin: 'localhost',
  })
)
app.use(urlencoded({ extended: false }))
app.use(json())

app.get('/health', (req: Request, res: Response) => {
  res.status(200).send('OK')
})

app.use('/push_notification', pushNotificationRoute)

app.use((err: Error, req: Request, res: Response, next: NextFunction) => {
  if (err) {
    logger.error(err)
    res.status(500).send(err.message)
  } else {
    next()
  }
})

app.use((req: Request, res: Response) => {
  res.status(404).send('Sorry cant find that')
})
app.disable('x-powered-by')

const PORT = config.app.port
app.listen(PORT, () => {
  logger.warn(`Server running on: ${PORT}`)
})

export default app
