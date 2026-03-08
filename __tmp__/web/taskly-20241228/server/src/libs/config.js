import 'dotenv/config';

const getEnv = (key, defaultValue) => process.env[key] || defaultValue;

export const Config = {
  APP_NAME: getEnv('APP_NAME', 'taskly-api'),
  APP_HOST: getEnv('APP_HOST', 'localhost'),
  APP_PORT: getEnv('APP_PORT', 8000),
  APP_LEVEL: getEnv('APP_LEVEL', 'development'),
  AUTH_SECRET: getEnv('AUTH_SECRET', 'auth-secret'),
  MONGO: {
    URI: getEnv('MONGO_URI', 'mongodb://admin:securepassword123@localhost:27017/'),
    DB: getEnv('MONGO_DB', 'taskly'),
    TIMEOUT: getEnv('MONGO_TIMEOUT', 30000),
  },
  CLIENT_URL: getEnv('CLIENT_URL', 'http://localhost:3000'),
};
