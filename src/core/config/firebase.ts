// import { SecretManagerServiceClient } from '@google-cloud/secret-manager'
import admin from 'firebase-admin'

// const getCredential = async (): Promise<string> => {
//   const client = new SecretManagerServiceClient()
//   const [accessResponse] = await client.accessSecretVersion({
//     name: 'FIREBASE_CREDENTIAL',
//   })
//   const responsePayload = accessResponse.payload?.data?.toString()
//   if (!responsePayload) {
//     throw new Error('getting erro in access secret manager')
//   }
//   return responsePayload
// }

import config from './enviroment_config'

admin.initializeApp({
  credential: admin.credential.applicationDefault(),
  databaseURL: config.firebase.databaseURL,
})

export const db = admin.firestore()
export const msg = admin.messaging()
