import express, { Request, Response } from 'express';
import { getPosts, getPost, createPost, updatePost, likePost, deletePost } from '../controllers/posts'; // Assuming the posts controllers are also written in TypeScript

const router = express.Router();

// Define routes with correct types for request and response
router.get('/', (req: Request, res: Response) => getPosts(req, res));
router.post('/', (req: Request, res: Response) => createPost(req, res));
router.get('/:id', (req: Request, res: Response) => getPost(req, res));
router.patch('/:id', (req: Request, res: Response) => updatePost(req, res));
router.delete('/:id', (req: Request, res: Response) => deletePost(req, res));
router.patch('/:id/likePost', (req: Request, res: Response) => likePost(req, res));

export default router;
