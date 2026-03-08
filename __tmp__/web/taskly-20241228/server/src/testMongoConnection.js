import { mongoDbInstance, closeMongo } from './libs/mongodb.js';
import log from './libs/logger.js';

async function testMongoConnection() {
  try {
    const db = mongoDbInstance;
    const collections = await db.listCollections().toArray();
    if (collections.length > 0) {
      log.info('MongoDB connection test succeeded!');
      log.info('Collections available:', collections);
    } else {
      log.warn('MongoDB connection test succeeded, but no collections found.');
    }
  } catch (error) {
    log.error('MongoDB connection test failed:', error);
  } finally {
    await closeMongo();
  }
}

testMongoConnection()
  .then(() => {
    log.info('MongoDB connection test completed.');
    process.exit(0); // Exit dengan kode sukses
  })
  .catch((err) => {
    log.error('MongoDB connection test failed.', err);
    process.exit(1); // Exit dengan kode gagal
  });
