import jwt from 'jsonwebtoken';
import log from '../libs/logger.js';
import { Config } from '../libs/config.js';

export const verifyToken = (req, res, next) => {
  try {
    const token = req.cookies.taskly_token;

    if (!token) {
      log.warn('Unauthorized access attempt: Token not provided');
      return next({ status: 401, message: 'Unauthorized: Token not provided' });
    }

    jwt.verify(token, Config.AUTH_SECRET, (err, user) => {
      if (err) {
        log.error(`JWT verification failed: ${err.message}`);
        return next({ status: 403, message: 'Forbidden: Invalid token' });
      }

      req.user = user;
      log.info(`Token verified for user: ${user.id}`);
      next();
    });
  } catch (error) {
    log.error('Unexpected error in verifyToken middleware', { error });
    next({ status: 500, message: 'Internal Server Error', error });
  }
};
