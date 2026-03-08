import express from 'express';
import authCtrl from '../controllers/auth.controller.js';
import userCtrl from '../controllers/user.controller.js';

const router = express.Router();

router.post('/api/v1/auth/signin', authCtrl.signin);
router.post('/api/v1/auth/signup', userCtrl.create);
router.post('/api/v1/auth/signout', authCtrl.signout);

export default router;
