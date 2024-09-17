import express, { Request, Response } from 'express';
import mongoose from 'mongoose';

import PostMessage from '../models/postMessage'; // Assuming this is already in TypeScript

const router = express.Router();

// Get all posts
export const getPosts = async (req: Request, res: Response): Promise<void> => { 
    try {
        const postMessages = await PostMessage.find();
        res.status(200).json(postMessages);
    } catch (error) {
        res.status(404).json({ message: (error as Error).message });
    }
};

// Get a single post by ID
export const getPost = async (req: Request, res: Response): Promise<void> => { 
    const { id } = req.params;

    try {
        const post = await PostMessage.findById(id);
        if (!post) return res.status(404).json({ message: "Post not found" });
        res.status(200).json(post);
    } catch (error) {
        res.status(404).json({ message: (error as Error).message });
    }
};

// Create a new post
export const createPost = async (req: Request, res: Response): Promise<void> => {
    const { title, message, selectedFile, creator, tags } = req.body;

    const newPostMessage = new PostMessage({ title, message, selectedFile, creator, tags });

    try {
        await newPostMessage.save();
        res.status(201).json(newPostMessage);
    } catch (error) {
        res.status(409).json({ message: (error as Error).message });
    }
};

// Update an existing post
export const updatePost = async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params;
    const { title, message, creator, selectedFile, tags } = req.body;
    
    if (!mongoose.Types.ObjectId.isValid(id)) return res.status(404).send(`No post with id: ${id}`);

    const updatedPost = { creator, title, message, tags, selectedFile, _id: id };

    try {
        const result = await PostMessage.findByIdAndUpdate(id, updatedPost, { new: true });
        if (!result) return res.status(404).json({ message: "Post not found" });
        res.json(result);
    } catch (error) {
        res.status(409).json({ message: (error as Error).message });
    }
};

// Delete a post
export const deletePost = async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params;

    if (!mongoose.Types.ObjectId.isValid(id)) return res.status(404).send(`No post with id: ${id}`);

    try {
        const result = await PostMessage.findByIdAndRemove(id);
        if (!result) return res.status(404).json({ message: "Post not found" });
        res.json({ message: "Post deleted successfully." });
    } catch (error) {
        res.status(409).json({ message: (error as Error).message });
    }
};

// Like a post
export const likePost = async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params;

    if (!mongoose.Types.ObjectId.isValid(id)) return res.status(404).send(`No post with id: ${id}`);
    
    try {
        const post = await PostMessage.findById(id);
        if (!post) return res.status(404).json({ message: "Post not found" });

        const updatedPost = await PostMessage.findByIdAndUpdate(id, { likeCount: post.likeCount + 1 }, { new: true });
        res.json(updatedPost);
    } catch (error) {
        res.status(409).json({ message: (error as Error).message });
    }
};

export default router;
