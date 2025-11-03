import express from 'express';
import upload from '../libs/local-upload.js';
import {
  getUsers,
  getUser,
  updateUser,
  deleteUser,
  uploadAvatar,
} from '../handlers/user-handler.js';
import { verifyToken } from '../middlewares/auth.js';

const router = express.Router();
router.get('/', getUsers);
router.get('/:id', getUser);
router.get('/:id', verifyToken, getUser);
router.patch('/update/:id', verifyToken, updateUser);
router.post('/upload/:id', verifyToken, upload.single('avatar'), uploadAvatar);
router.delete('/delete/:id', verifyToken, deleteUser);

export default router;
