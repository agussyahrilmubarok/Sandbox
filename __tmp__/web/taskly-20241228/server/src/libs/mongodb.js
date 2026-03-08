import { MongoClient } from 'mongodb';
import { Config } from './config.js';
import log from './logger.js';

let mongoDb = null;
const url = Config.MONGO.URI;
const dbName = Config.MONGO.DB;
const client = new MongoClient(url);

export const connectMongo = async () => {
  if (mongoDb) {
    log.info('MongoDB connection already established.');
    return mongoDb;
  }

  try {
    await client.connect();
    log.info('MongoDB connected successfully.');
    mongoDb = client.db(dbName);
    return mongoDb;
  } catch (error) {
    log.error('Failed to connect to MongoDB.', error);
    throw error;
  }
};

export const getDb = () => {
  if (!mongoDb) {
    throw new Error('MongoDB not connected.');
  }
  return mongoDb;
};

export const closeMongo = async () => {
  if (mongoDb) {
    await client.close();
    log.info('MongoDB connection closed.');
    mongoDb = null;
  } else {
    log.info('No active MongoDB connection to close.');
  }
};

export const mongoDbInstance = await connectMongo();
