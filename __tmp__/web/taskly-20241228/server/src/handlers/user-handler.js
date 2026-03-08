import bcrypt from 'bcrypt';
import { ObjectId } from 'mongodb';

import log from '../libs/logger.js';
import { mongoDbInstance } from '../libs/mongodb.js';

// MongoDB user collection
const userCollection = mongoDbInstance.collection('users');

/**
 * Fetch a list of users based on query filters (username or email).
 * @param {Object} req - HTTP request object.
 * @param {Object} res - HTTP response object.
 * @param {Function} next - Middleware function for error handling.
 */
export const getUsers = async (req, res, next) => {
  try {
    const { username, email } = req.query;
    let query = {};

    if (username) {
      query.username = { $regex: new RegExp(username, 'i') }; // Case-insensitive search
    }

    if (email) {
      query.email = { $regex: new RegExp(email, 'i') }; // Case-insensitive search
    }

    log.info('Fetching users with query:', query);
    const users = await userCollection.find(query).toArray();

    if (!users.length) {
      log.warn('No users found with the given query filters');
      return next({ status: 404, message: 'No users found' });
    }

    // Map users to include avatar URLs
    const rest = users.map((user) => {
      const { password, ...u } = user;
      return {
        ...u,
      };
    });

    res.status(200).json(rest);
  } catch (error) {
    log.error('Error fetching users:', error);
    next({ status: 500, message: 'Internal server error' });
  }
};

/**
 * Fetch a single user by their ID.
 * @param {Object} req - HTTP request object.
 * @param {Object} res - HTTP response object.
 * @param {Function} next - Middleware function for error handling.
 */
export const getUser = async (req, res, next) => {
  try {
    const userId = req.params.id;
    if (!ObjectId.isValid(userId)) {
      return next({ status: 400, message: 'Invalid user ID format' });
    }

    log.info('Fetching user with ID:', userId);
    const query = { _id: new ObjectId(userId) };
    const user = await userCollection.findOne(query);

    if (!user) {
      log.warn('User not found with ID:', userId);
      return next({ status: 404, message: 'User not found!' });
    }

    log.info('Successfully fetched user:', user);
    res.status(200).json({
      ...user,
      password: undefined,
    });
  } catch (error) {
    log.error('Error fetching user:', error);
    next({ status: 500, message: 'Internal server error' });
  }
};

/**
 * Update user information by their ID.
 * @param {Object} req - HTTP request object.
 * @param {Object} res - HTTP response object.
 * @param {Function} next - Middleware function for error handling.
 */
export const updateUser = async (req, res, next) => {
  // Validate if user is trying to update their own account
  if (req.user.id !== req.params.id) {
    log.warn('Unauthorized update attempt by user:', req.user.id);
    return next({
      status: 401,
      message: 'You can only update your own account',
    });
  }

  try {
    const userId = req.params.id;
    if (!ObjectId.isValid(userId)) {
      log.warn('Invalid user ID format:', userId);
      return next({ status: 400, message: 'Invalid user ID format' });
    }

    // Check if password is provided, hash it
    if (req.body.password) {
      req.body.password = await bcrypt.hash(req.body.password, 10);
    }

    // Prepare update data
    const query = { _id: new ObjectId(userId) };
    const data = {
      $set: {
        ...req.body,
        updatedAt: new Date().toISOString(),
      },
    };
    const options = { returnDocument: 'after' };

    const updatedUser = await userCollection.findOneAndUpdate(
      query,
      data,
      options
    );
    if (!updatedUser) {
      log.warn('User not found for update:', userId);
      return next({ status: 404, message: 'User not found!' });
    }

    log.info('Successfully updated user:', updatedUser);
    res.status(200).json({
      ...updatedUser,
      password: undefined,
    });
  } catch (error) {
    log.error('Error updating user:', error);
    next({ status: 500, message: 'Internal server error' });
  }
};

/**
 * Delete a user by their ID.
 * @param {Object} req - HTTP request object.
 * @param {Object} res - HTTP response object.
 * @param {Function} next - Middleware function for error handling.
 */
export const deleteUser = async (req, res, next) => {
  // Validate if user is trying to delete their own account
  if (req.user.id !== req.params.id) {
    log.warn('Unauthorized delete attempt by user:', req.user.id);
    return next({
      status: 401,
      message: 'You can only delete your own account',
    });
  }

  try {
    const userId = req.params.id;
    if (!ObjectId.isValid(userId)) {
      log.warn('Invalid user ID format:', userId);
      return next({ status: 400, message: 'Invalid user ID format' });
    }

    const query = { _id: new ObjectId(userId) };
    const deleteResult = await userCollection.deleteOne(query);
    if (deleteResult.deletedCount === 0) {
      log.warn('Invalid user ID format:', userId);
      return next({ status: 404, message: 'User not found!' });
    }

    log.info('Successfully deleted user:', userId);
    res.status(200).json({ message: 'User has been deleted' });
  } catch (error) {
    log.error('Error deleting user:', error);
    next({ status: 500, message: 'Internal server error' });
  }
};

/**
 * Upload a new avatar for the user.
 * @param {Object} req - HTTP request object containing the file and user details.
 * @param {Object} res - HTTP response object.
 * @param {Function} next - Middleware function for error handling.
 */
export const uploadAvatar = async (req, res, next) => {
  if (req.user.id !== req.params.id) {
    log.warn(
      `Unauthorized avatar upload attempt by user: ${req.user.id}, target user: ${req.params.id}`
    );
    return next({
      status: 401,
      message: 'You can only upload an avatar for your own account',
    });
  }

  try {
    // Check if a file is uploaded
    if (!req.file) {
      log.warn('No file uploaded in the request');
      return res.status(400).json({ message: 'No file uploaded' });
    }

    // Validate file size (e.g., max 2MB)
    if (req.file.size > 2 * 1024 * 1024) {
      log.warn(
        `Uploaded file exceeds size limit. File size: ${req.file.size} bytes`
      );
      return res.status(400).json({ message: 'File size exceeds 2MB' });
    }

    // Get the user ID from the URL params
    const userId = req.user.id;
    if (!ObjectId.isValid(userId)) {
      log.warn(`Invalid user ID format during avatar upload: ${userId}`);
      return next({ status: 400, message: 'Invalid user ID format' });
    }

    const query = { _id: new ObjectId(userId) };
    const updateData = {
      $set: {
        avatar: `/images/users/${req.file.filename}`, // Save the filename
        updatedAt: new Date().toISOString(),
      },
    };
    const options = { returnDocument: 'after' };

    const updatedUser = await userCollection.findOneAndUpdate(
      query,
      updateData,
      options
    );
    if (!updatedUser) {
      log.warn(`User not found for avatar update: ${userId}`);
      return next({ status: 404, message: 'User not found!' });
    }

    res.status(200).json({
      ...updatedUser,
      password: undefined,
    });
  } catch (error) {
    // Handle any errors
    log.error('Error uploading avatar:', error);
    return res.status(500).json({ message: 'Internal server error' });
  }
};
