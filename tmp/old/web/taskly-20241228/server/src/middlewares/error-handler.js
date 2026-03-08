import log from '../libs/logger.js';

export const errorHandler = (err, req, res, next) => {
  const defaultMessage =
    "We're having technical issues. Please try again later.";

  // Log error details
  log.error('Error occurred', {
    message: err.message,
    status: err.status,
    stack: err.error?.stack || err.stack || 'No stack available',
  });

  // Send structured error response
  res.status(err.status || 500).json({
    status: err.status || 500,
    message: err.message || defaultMessage,
  });
};
