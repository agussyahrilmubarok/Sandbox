import express from 'express';
import {
  getTask,
  getTasksByUser,
  createTask,
  updateTask,
  deleteTask,
} from '../handlers/task-handler.js';
import { verifyToken } from '../middlewares/auth.js';

const router = express.Router();
router.get('/:id', verifyToken, getTask);
router.get('/user/:id', verifyToken, getTasksByUser);
router.post('/create', verifyToken, createTask);
router.patch('/:id', verifyToken, updateTask);
router.delete('/:id', verifyToken, deleteTask);

export default router;
