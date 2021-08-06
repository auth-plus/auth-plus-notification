declare global {
  namespace NodeJS {
    interface ProcessEnv {
      NODE_ENV: 'development' | 'production' | 'test'
      APP_PORT: string
      APP_NAME: string
      DATABASE_HOST: string
      DATABASE_USER: string
      DATABASE_PASSWORD: string
      DATABASE_DATABASE: string
      DATABASE_PORT: string
      DATABASE_URL: string
      FIREBASE_DATABASE_URL: string
    }
  }
}

export {}
