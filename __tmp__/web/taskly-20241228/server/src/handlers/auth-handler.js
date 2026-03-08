import bcrypt from 'bcrypt';
import jwt from 'jsonwebtoken';

import log from '../libs/logger.js';
import { Config } from '../libs/config.js';
import { mongoDbInstance } from '../libs/mongodb.js';
import { responseWithCookie } from '../libs/response-resolver.js';

// MongoDB user collection
const userCollection = mongoDbInstance.collection('users');

/**
 * Handle user signup process, including validation, password hashing,
 * storing user data in the database, and generating a JWT token for authentication.
 *
 * @param {Object} req - The request object containing the user details from the client.
 * @param {Object} res - The response object to send back the result to the client.
 * @param {Function} next - The next middleware function to call if an error occurs.
 */
export const signUp = async (req, res, next) => {
  try {
    const { username, email, password } = req.body;

    const query = { $or: [{ email }, { username }] };
    const existingUser = await userCollection.findOne(query);
    if (existingUser) {
      log.warn(
        `Signup attempt failed: ${email} or ${username} already registered`
      );
      return next({
        status: 422,
        message: 'Email or Username is already registered.',
      });
    }

    // Hash the password securely before storing it in the database
    const hashedPassword = await bcrypt.hash(password, 10);
    const user = {
      username,
      email,
      password: hashedPassword,
      avatar: '/images/users/default-user.png', // Default avatar for the new user
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };

    const { insertedId } = await userCollection.insertOne(user);
    const token = jwt.sign({ id: insertedId }, Config.AUTH_SECRET);
    const { password: pass, ...rest } = user;

    log.info(`New user signed up successfully: ${username} (${email})`);
    responseWithCookie(res, token, { ...rest });
  } catch (error) {
    log.error('Error during signup process:', error);
    next({ status: 500, error });
  }
};

/**
 * Handle user sign-in process by verifying the provided credentials,
 * generating a JWT token for the session, and responding with user data.
 *
 * @param {Object} req - The request object containing user credentials.
 * @param {Object} res - The response object to send back the result.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const signIn = async (req, res, next) => {
  const { email, password } = req.body;

  try {
    const validUser = await userCollection.findOne({ email });
    if (!validUser) {
      log.warn(`Sign-in failed: User not found (${email})`);
      return next({ status: 404, message: 'User not found!' });
    }

    // Verify password by comparing it with the stored hash
    const validPassword = await bcrypt.compare(password, validUser.password);
    if (!validPassword) {
      log.warn(`Sign-in failed: Invalid password for user (${email})`);
      return next({ status: 401, message: 'Wrong password!' });
    }

    const token = jwt.sign({ id: validUser._id }, Config.APP_SECRET);
    const { password: pass, ...rest } = validUser;

    log.info(`User signed in successfully: ${validUser.username} (${email})`);
    responseWithCookie(res, token, { ...rest });
  } catch (error) {
    log.error('Error during sign-in process:', error);
    next({ status: 500, error });
  }
};

/**
 * Handle user sign-out process, which clears the authentication token from the user's session.
 *
 * @param {Object} req - The request object (does not require credentials).
 * @param {Object} res - The response object to send the result.
 * @param {Function} next - The next middleware function to call in case of an error.
 */
export const signOut = async (req, res, next) => {
  try {
    log.info(`User signed out successfully`);

    res.clearCookie('taskly_token');

    res.status(200).json({ message: 'Sign out successful' });
  } catch (error) {
    log.error('Error during sign-out process:', error);
    next({ status: 500, error });
  }
};
