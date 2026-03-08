import bcrypt from 'bcrypt';
import { ObjectId } from 'mongodb';

/**
 * User Model
 * This class represents a user entity.
 */
class User {
  constructor(username, email, password, avatar = 'default-user.png') {
    this.username = username;
    this.email = email;
    this.password = password;
    this.avatar = avatar;
    this.createdAt = new Date().toISOString();
    this.updatedAt = new Date().toISOString();
  }

  // Method to hash the password before saving it to the database
  async hashPassword() {
    this.password = await bcrypt.hash(this.password, 10);
  }

  // Static method to validate the password
  static async validatePassword(inputPassword, hashedPassword) {
    return bcrypt.compare(inputPassword, hashedPassword);
  }

  // Static method to map MongoDB's _id field to custom id field
  static mapMongoIdToCustomId(user) {
    if (user && user._id) {
      user.id = user._id.toString(); // Map _id to id
      delete user._id; // Remove _id from the returned object
    }
    return user;
  }

  // Static method to convert DB object into model object
  static fromDbObject(dbObject) {
    const user = new User(
      dbObject.username,
      dbObject.email,
      dbObject.password,
      dbObject.avatar
    );
    user.createdAt = dbObject.createdAt;
    user.updatedAt = dbObject.updatedAt;
    user.id = dbObject._id ? dbObject._id.toString() : dbObject.id;
    return user;
  }

  // Method to convert model object into DB object
  toDbObject() {
    return {
      username: this.username,
      email: this.email,
      password: this.password,
      avatar: this.avatar,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }

  // Static method to get the schema for MongoDB or SQL usage
  static getSchema() {
    return {
      id: { type: 'string', required: true },
      username: { type: 'string', required: true },
      email: { type: 'string', required: true },
      password: { type: 'string', required: true },
      avatar: { type: 'string', default: 'default-user.png' },
      createdAt: { type: 'string', required: true },
      updatedAt: { type: 'string', required: true },
    };
  }

  // Method to return a user object formatted for MongoDB or any other DB
  static toDbFormat(user) {
    const dbObject = user.toDbObject();
    if (user.id) {
      dbObject._id = new ObjectId(user.id); // For MongoDB _id
    }
    return dbObject;
  }

  // Static method to convert a DB object into a user instance
  static fromDbFormat(dbObject) {
    return User.fromDbObject(dbObject);
  }
}

export default User;
