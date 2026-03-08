import { ObjectId } from 'mongodb';

import log from '../libs/logger.js';
import { mongoDbInstance } from '../libs/mongodb.js';

// MongoDB task collection
const taskCollection = mongoDbInstance.collection('tasks');

/**
 * Get tasks owned by a specific user with optional filters for status and ordering.
 * Supports pagination and sorting.
 *
 * @param {Object} req - The request object containing parameters and query parameters.
 * @param {Object} res - The response object to send back the result to the client.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const getTasksByUser = async (req, res, next) => {
  try {
    const query = { owner: new ObjectId(req.params.id) }; // Ensure owner is an ObjectId
    const { status, orderBy } = req.query;
    const sort = orderBy ? { [orderBy]: 1 } : {}; // Sort by the given field

    // If status filter is provided, add it to the query
    if (status) {
      query['status'] = status;
    }

    // Pagination setup
    const page = parseInt(req.query.page) || 1;
    const pageSize = 4;

    // Fetch tasks with pagination and sorting
    const tasks = await taskCollection
      .find(query)
      .sort(sort)
      .limit(pageSize)
      .skip((page - 1) * pageSize)
      .toArray();

    // Get total task count for pagination
    const taskCount = await taskCollection.countDocuments(query);

    log.info(`Fetched tasks for user with ID: ${req.params.id}`, {
      tasks,
      taskCount,
    });

    // Return tasks and task count
    res.status(200).json({ tasks, taskCount });
  } catch (error) {
    log.error('Error fetching tasks by user:', error);
    next({ status: 500, error });
  }
};

/**
 * Get a specific task by its ID.
 *
 * @param {Object} req - The request object containing the task ID.
 * @param {Object} res - The response object to send back the task details.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const getTask = async (req, res, next) => {
  try {
    const taskId = req.params.id;
    const userId = req.user.id;

    const task = await taskCollection.findOne({ _id: new ObjectId(taskId) });
    if (!task) {
      log.warn(`Task not found: ${taskId}`);
      return next({ status: 404, message: 'Task not found!' });
    }

    if (task.owner.toString() !== userId) {
      log.warn(
        `Unauthorized update attempt by user ${userId} on task ${taskId}`
      );
      return next({
        status: 403,
        message: 'You are not authorized to update this task',
      });
    }

    log.info(`Fetched task: ${req.params.id}`, { task });

    // Return the task
    res.status(200).json(task);
  } catch (error) {
    log.error('Error fetching task:', error);
    next({ status: 500, error });
  }
};

/**
 * Create a new task for the authenticated user.
 *
 * @param {Object} req - The request object containing task details.
 * @param {Object} res - The response object to send back the result.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const createTask = async (req, res, next) => {
  try {
    const newTask = req.body;
    newTask.owner = new ObjectId(req.user.id); // Ensure owner is an ObjectId
    newTask.createdAt = new Date().toISOString();
    newTask.updatedAt = new Date().toISOString();

    // Insert the new task into the database
    const result = await taskCollection.insertOne(newTask);
    log.info('Task created successfully', {
      taskId: result.insertedId,
      owner: req.user.id,
    });

    // Fetch the inserted task by its insertedId
    const createdTask = await taskCollection.findOne({
      _id: result.insertedId,
    });

    // Return the created task
    res.status(200).json(createdTask); // Return the inserted task
  } catch (error) {
    log.error('Error creating task:', error);
    next({ status: 500, error });
  }
};

/**
 * Update an existing task.
 *
 * @param {Object} req - The request object containing the task ID and updated task data.
 * @param {Object} res - The response object to send back the updated task.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const updateTask = async (req, res, next) => {
  try {
    const taskId = req.params.id;
    const userId = req.user.id;

    const task = await taskCollection.findOne({ _id: new ObjectId(taskId) });
    if (!task) {
      log.warn(`Task not found: ${taskId}`);
      return next({ status: 404, message: 'Task not found!' });
    }

    if (task.owner.toString() !== userId) {
      log.warn(
        `Unauthorized update attempt by user ${userId} on task ${taskId}`
      );
      return next({
        status: 403,
        message: 'You are not authorized to update this task',
      });
    }

    const data = {
      $set: {
        ...req.body, // Menyertakan data yang diupdate
        updatedAt: new Date().toISOString(), // Memperbarui waktu update
      },
    };

    const updatedTask = await taskCollection.findOneAndUpdate(
      { _id: new ObjectId(taskId) },
      data,
      { returnDocument: 'after' } // Mengembalikan document setelah update
    );

    if (!updatedTask) {
      log.warn(`Task update failed, task not found: ${taskId}`);
      return next({ status: 404, message: 'Task not found!' });
    }

    log.info(`Task updated successfully: ${taskId}`, {
      updatedTask: updatedTask,
    });

    // Return the updated task
    res.status(200).json(updatedTask);
  } catch (error) {
    log.error('Error updating task:', error);
    next({ status: 500, error });
  }
};

/**
 * Delete a task by its ID.
 *
 * @param {Object} req - The request object containing the task ID.
 * @param {Object} res - The response object to send the result.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const deleteTask = async (req, res, next) => {
  try {
    const taskId = req.params.id;
    const userId = req.user.id;

    const task = await taskCollection.findOne({ _id: new ObjectId(taskId) });
    if (!task) {
      log.warn(`Task not found: ${taskId}`);
      return next({ status: 404, message: 'Task not found!' });
    }

    if (task.owner.toString() !== userId) {
      log.warn(
        `Unauthorized update attempt by user ${userId} on task ${taskId}`
      );
      return next({
        status: 403,
        message: 'You are not authorized to update this task',
      });
    }

    const result = await taskCollection.deleteOne({
      _id: new ObjectId(taskId),
    });

    if (result.deletedCount === 0) {
      log.warn(`Task delete failed, task not found: ${req.params.id}`);
      return next({ status: 404, message: 'Task not found!' });
    }

    log.info(`Task deleted successfully: ${req.params.id}`);

    // Return success message
    res.status(200).json({ message: 'Task has been deleted' });
  } catch (error) {
    log.error('Error deleting task:', error);
    next({ status: 500, error });
  }
};
