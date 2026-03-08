import express from 'express';
import userCtrl from '../controllers/user.controller.js';
import authCtrl from '../controllers/auth.controller.js';

const router = express.Router();

// 🔹 Public routes
router.route('/api/v1/users')
    .get(userCtrl.list);

// 🔹 Protected routes (require authentication)
router.route('/api/v1/users/:userId')
    .get(authCtrl.requireSignin, userCtrl.read)
    .put(authCtrl.requireSignin, authCtrl.hasAuthorization, userCtrl.update)
    .delete(authCtrl.requireSignin, authCtrl.hasAuthorization, userCtrl.remove);

// 🔹 Middleware to load user by ID when `userId` param is present
router.param('userId', userCtrl.userByID);

export default router;
