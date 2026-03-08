import { mongoDbInstance, closeMongo } from './libs/mongodb.js';
import bcrypt from 'bcrypt';
import log from './libs/logger.js';

const users = [
  {
    username: 'johndoe123',
    email: 'johndoe123@mail.com',
    password: 'password123',
    avatar: '/images/users/default-avatar.jpg',
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
  {
    username: 'jane78',
    email: 'jane@mail.com',
    password: 'securepassword456',
    avatar: '/images/users/default-avatar.jpg',
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
];

const tasks = [
  {
    name: 'Read Atomic Habits',
    description: 'Finish reading Atomic Habits by James Clear',
    priority: 'not urgent',
    due: new Date().toISOString(),
    status: 'open',
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
  {
    name: 'Learn MERN Stack',
    description:
      'Learn the MERN stack and build a full-stack application with it',
    priority: 'urgent',
    due: new Date().toISOString(),
    status: 'open',
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
];

const getEmotion = (level) => {
  switch (level) {
    case 'info':
      return '😊';
    case 'warn':
      return '😟';
    case 'error':
      return '😢';
    default:
      return '⚡';
  }
};

async function hashPasswords(users) {
  try {
    for (let user of users) {
      user.password = await bcrypt.hash(user.password, 10);
    }
    log.info(`${getEmotion('info')} Passwords have been hashed successfully.`);
  } catch (error) {
    log.error(`${getEmotion('error')} Error hashing passwords: ${error}`);
    throw error;
  }
}

async function checkIfUsersExist() {
  const usersCollection = mongoDbInstance.collection('users');
  const existingUsers = await usersCollection
    .find({
      $or: [
        { email: { $in: users.map((user) => user.email) } },
        { username: { $in: users.map((user) => user.username) } },
      ],
    })
    .toArray();
  return existingUsers;
}

async function checkIfTasksExist() {
  const tasksCollection = mongoDbInstance.collection('tasks');
  const taskNames = tasks.map((task) => task.name);
  const existingTasks = await tasksCollection
    .find({
      name: { $in: taskNames },
    })
    .toArray();
  return existingTasks;
}

async function insertUsers(users) {
  try {
    const usersCollection = mongoDbInstance.collection('users');
    const userInsertResult = await usersCollection.insertMany(users);
    log.info(
      `${getEmotion('info')} Inserted ${userInsertResult.insertedCount} users into the database.`
    );
    return userInsertResult.insertedIds;
  } catch (error) {
    log.error(
      `${getEmotion('error')} Error inserting users into the database: ${error}`
    );
    throw error;
  }
}

async function insertTasks(tasks, userIds) {
  try {
    tasks[0].owner = userIds[0]; // Assign task to first user
    tasks[1].owner = userIds[1]; // Assign task to second user
    const tasksCollection = mongoDbInstance.collection('tasks');
    await tasksCollection.insertMany(tasks);
    log.info(
      `${getEmotion('info')} Inserted ${tasks.length} tasks into the database.`
    );
  } catch (error) {
    log.error(
      `${getEmotion('error')} Error inserting tasks into the database: ${error}`
    );
    throw error;
  }
}

async function seedData() {
  try {
    log.info(`${getEmotion('info')} Seeding database started...`);
    const existingUsers = await checkIfUsersExist();
    if (existingUsers.length > 0) {
      log.info(
        `${getEmotion('warn')} Users already exist. Skipping user seeding.`
      );
    } else {
      await hashPasswords(users);
      const userIds = await insertUsers(users);
      await insertTasks(tasks, userIds);
      log.info(`${getEmotion('info')} Data seeding completed successfully.`);
    }
  } catch (error) {
    log.error(`${getEmotion('error')} Data seeding failed: ${error}`);
    throw error; // Ensure the process halts on failure
  } finally {
    await closeMongo();
  }
}

seedData()
  .then(() => {
    log.info(`${getEmotion('info')} Seeder completed successfully.`);
    process.exit(0);
  })
  .catch((err) => {
    log.error(`${getEmotion('error')} Seeder failed. ${err}`);
    process.exit(1);
  });
