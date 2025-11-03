import express from 'express';
import cookieParser from 'cookie-parser';
import cors from 'cors';

import log from './libs/logger.js';
import authRoutes from './routes/auth-route.js';
import userRoutes from './routes/user-route.js';
import taskRoutes from './routes/task-route.js';

import { Config } from './libs/config.js';
import { closeMongo } from './libs/mongodb.js';
import { errorHandler } from './middlewares/error-handler.js';

const app = express();
const port = Config.APP_PORT;
const host = Config.APP_HOST;

app.use(
  cors({
    origin: Config.CLIENT_URL,
    credentials: true,
    methods: 'GET, POST, PUT, DELETE, OPTIONS, PATCH',
    allowedHeaders: 'Content-Type, Authorization',
  })
);

app.use(express.json());
app.use(cookieParser());
app.use(express.static('public'));

// Routes
app.get('/', (req, res) => {
  log.info('GET / request received');
  log.info('Cookies: ', req.cookies);
  res.status(200).json({ message: 'Welcome to the Taskly API' });
});

// Authentication and Authorization routes
app.use('/api/v1/auth', authRoutes);

// Users api features
app.use('/api/v1/users', userRoutes);

// Tasks user api features
app.use('/api/v1/tasks', taskRoutes);

// Catch-all for undefined routes
app.use('*', (req, res) => {
  log.warn(`404 not found: ${req.originalUrl}`);
  res.status(404).json({ message: 'Not found' });
});

// Error handler middleware
app.use(errorHandler);

// Start Express server
app.listen(port, () => {
  log.info(`Server is running on http://${host}:${port}`);
});

// Graceful shutdown function
const gracefulShutdown = async () => {
  log.info('Server is shutting down gracefully...');

  // Close MongoDB connection
  closeMongo()
    .then(() => {
      log.info('MongoDB connection closed.');
      process.exit(0); // Graceful exit
    })
    .catch((error) => {
      log.error('Error closing MongoDB connection', error);
      process.exit(1); // Force exit on error
    });
};

// Handling SIGINT and SIGTERM signals for graceful shutdown
process.on('SIGINT', gracefulShutdown); // CTRL+C
process.on('SIGTERM', gracefulShutdown); // SIGTERM (e.g., from PM2 or Docker)
