import User from '../models/user.model.js';
import errorHandler from '../helpers/dbErrorHandler.js';

/** 🔹 Create a new user */
const create = async (req, res) => {
    try {
        const user = new User(req.body);
        await user.save();
        return res.status(201).json({ message: "Successfully signed up!" });
    } catch (err) {
        return res.status(400).json({ error: errorHandler.getErrorMessage(err) || err.message });
    }
};

/** 🔹 Middleware: Load user by ID */
const userByID = async (req, res, next, id) => {
    try {
        const user = await User.findById(id);
        if (!user) return res.status(404).json({ error: "User not found" });

        req.profile = user;
        next();
    } catch (err) {
        return res.status(400).json({ error: "Could not retrieve user" });
    }
};

/** 🔹 Get a single user */
const read = (req, res) => {
    const { _id, name, email, createdAt, updatedAt } = req.profile;
    return res.json({ _id, name, email, createdAt, updatedAt });
};

/** 🔹 List all users */
const list = async (req, res) => {
    try {
        const users = await User.find().select('_id name email createdAt updatedAt');
        res.json(users);
    } catch (err) {
        return res.status(400).json({ error: errorHandler.getErrorMessage(err) || err.message });
    }
};

/** 🔹 Update user */
const update = async (req, res) => {
    try {
        const { name, email } = req.body;
        const updatedUser = await User.findByIdAndUpdate(req.profile._id, { name, email, updatedAt: Date.now() }, { new: true, runValidators: true })
            .select('_id name email createdAt updatedAt'); // Prevent returning hashed_password

        res.json(updatedUser);
    } catch (err) {
        return res.status(400).json({ error: errorHandler.getErrorMessage(err) || err.message });
    }
};

/** 🔹 Delete user */
const remove = async (req, res) => {
    try {
        const deletedUser = await User.findByIdAndDelete(req.profile._id).select('_id name email');
        if (!deletedUser) return res.status(404).json({ error: "User not found" });

        res.json({ message: "User deleted successfully", deletedUser });
    } catch (err) {
        return res.status(400).json({ error: errorHandler.getErrorMessage(err) || err.message });
    }
};

export default { create, userByID, read, list, update, remove };
