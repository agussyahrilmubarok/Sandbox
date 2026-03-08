import { ObjectId } from 'mongodb';

/**
 * Task Model
 * This class represents a task entity.
 */
class Task {
  constructor(name, description, priority, due, status = 'open', owner) {
    this.name = name;
    this.description = description;
    this.priority = priority;
    this.due = due;
    this.status = status;
    this.owner = owner; // This would be a reference to a User
    this.createdAt = new Date().toISOString();
    this.updatedAt = new Date().toISOString();
  }

  // Static method to map MongoDB's _id field to custom id field
  static mapMongoIdToCustomId(task) {
    if (task && task._id) {
      task.id = task._id.toString(); // Map _id to id
      delete task._id; // Remove _id from the returned object
    }
    return task;
  }

  // Static method to convert DB object into model object
  static fromDbObject(dbObject) {
    const task = new Task(
      dbObject.name,
      dbObject.description,
      dbObject.priority,
      dbObject.due,
      dbObject.status,
      dbObject.owner
    );
    task.createdAt = dbObject.createdAt;
    task.updatedAt = dbObject.updatedAt;
    task.id = dbObject._id ? dbObject._id.toString() : dbObject.id;
    return task;
  }

  // Method to convert model object into DB object
  toDbObject() {
    return {
      name: this.name,
      description: this.description,
      priority: this.priority,
      due: this.due,
      status: this.status,
      owner: this.owner,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }

  // Static method to get the schema for MongoDB or SQL usage
  static getSchema() {
    return {
      id: { type: 'string', required: true },
      name: { type: 'string', required: true },
      description: { type: 'string', required: true },
      priority: { type: 'string', required: true },
      due: { type: 'string', required: true },
      status: { type: 'string', required: true },
      owner: { type: 'string', required: true }, // Reference to a user
      createdAt: { type: 'string', required: true },
      updatedAt: { type: 'string', required: true },
    };
  }

  // Method to return a task object formatted for MongoDB or any other DB
  static toDbFormat(task) {
    const dbObject = task.toDbObject();
    if (task.id) {
      dbObject._id = new ObjectId(task.id); // For MongoDB _id
    }
    return dbObject;
  }

  // Static method to convert a DB object into a task instance
  static fromDbFormat(dbObject) {
    return Task.fromDbObject(dbObject);
  }
}

export default Task;
