import winston from 'winston';
import DailyRotateFile from 'winston-daily-rotate-file';

import { Config } from './config.js';

const serviceName = Config.APP_NAME;
const isProduction = Config.APP_LEVEL == 'production' ? true : false;

const dailyRotateTransport = new DailyRotateFile({
  dirname: './logs',
  filename: 'taskly-api-%DATE%.log',
  datePattern: 'YYYY-MM-DD',
  zippedArchive: true,
  maxSize: '20m',
  maxFiles: '14d',
});

const log = winston.createLogger({
  level: isProduction ? 'warn' : 'debug',
  defaultMeta: {
    service: serviceName,
  },
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.json()
  ),
  transports: [
    new winston.transports.Console({
      format: isProduction
        ? winston.format.combine(
            winston.format.timestamp(),
            winston.format.simple()
          )
        : winston.format.combine(
            winston.format.colorize(),
            winston.format.simple()
          ),
    }),
    dailyRotateTransport,
  ],
});

export default log;
