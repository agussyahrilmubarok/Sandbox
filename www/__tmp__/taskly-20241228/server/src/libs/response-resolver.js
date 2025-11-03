import { Config } from './config.js';

export const responseWithCookie = (res, token, rest) => {
  if (Config.APP_LEVEL === 'production') {
    res
      .cookie('taskly_token', token, {
        httpOnly: true, // Cannot be accessed by JavaScript
        sameSite: 'None', // SameSite None is required for cross-origin requests
        secure: true, // Only sent over HTTPS
        maxAge: 60 * 60 * 1000, // Cookie expiration time, e.g., 1 hour
        domain: Config.CLIENT_URL, // For example, .example.com to support subdomains
        path: '/', // Cookie applies across the whole domain
      })
      .status(200)
      .json(rest);
  } else {
    res
      .cookie('taskly_token', token, {
        httpOnly: true, // Cannot be accessed by JavaScript
        sameSite: 'Lax', // SameSite Lax is safer in development
        secure: false, // No need for secure in localhost (since it's not HTTPS)
        maxAge: 60 * 60 * 1000, // Cookie expiration time, e.g., 1 hour
        path: '/', // Cookie applies across the whole domain
      })
      .status(200)
      .json(rest);
  }
};
