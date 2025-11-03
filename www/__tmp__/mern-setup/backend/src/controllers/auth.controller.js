import User from '../models/user.model.js';
import jwt from 'jsonwebtoken';
import config from '../config.js';
import bcrypt from 'bcryptjs';

// 🔹 SIGNIN (Login)
const signin = async (req, res) => {
    try {
        const { email, password } = req.body;

        let user = await User.findOne({ email });
        if (!user) {
            return res.status(401).json({ error: "User not found" });
        }

        // Check password using bcryptjs
        const isMatch = await bcrypt.compare(password, user.password);
        if (!isMatch) {
            return res.status(401).json({ error: "Email and password don't match." });
        }

        // Generate JWT token
        const token = jwt.sign({ _id: user._id }, config.jwtSecret, {
            expiresIn: "7d", // Token expires in 7 days
        });

        // Store token in an HTTP-only cookie
        res.cookie("t", token, { httpOnly: true, maxAge: 7 * 24 * 60 * 60 * 1000 });

        return res.json({
            token,
            user: {
                _id: user._id,
                name: user.name,
                email: user.email,
            },
        });

    } catch (err) {
        return res.status(500).json({ error: "Could not sign in" });
    }
};

// 🔹 SIGNOUT (Logout)
const signout = (req, res) => {
    res.clearCookie("t");
    return res.status(200).json({ message: "Signed out successfully" });
};

// 🔹 Middleware: Verify JWT Token
const requireSignin = (req, res, next) => {
    try {
        const token = req.cookies.t || req.headers.authorization?.split(" ")[1];
        if (!token) {
            return res.status(401).json({ error: "Access denied. No token provided." });
        }

        const decoded = jwt.verify(token, config.jwtSecret);
        req.auth = decoded;
        next();
    } catch (err) {
        return res.status(401).json({ error: "Invalid or expired token" });
    }
};

// 🔹 Middleware: Check User Authorization
const hasAuthorization = (req, res, next) => {
    if (!req.profile || req.auth._id !== req.profile._id.toString()) {
        return res.status(403).json({ error: "User is not authorized" });
    }
    next();
};

export default {
    signin,
    signout,
    requireSignin,
    hasAuthorization,
};
