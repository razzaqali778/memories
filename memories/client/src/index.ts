import express, { Application, Request, Response, NextFunction } from 'express';
import bodyParser from 'body-parser';
import mongoose from 'mongoose';
import cors from 'cors';

import postRoutes from './routes/posts.js'; // Ensure this file is in TypeScript or compiled

const app: Application = express();

// Middlewares
app.use(bodyParser.json({ limit: '30mb', extended: true }));
app.use(bodyParser.urlencoded({ limit: '30mb', extended: true }));
app.use(cors());

// Routes
app.use('/posts', postRoutes);

// MongoDB connection URL and Server port
const CONNECTION_URL: string = 'mongodb+srv://js_mastery:123123123@practice.jto9p.mongodb.net/test';
const PORT: string | number = process.env.PORT || 5000;

// Mongoose connection
mongoose.connect(CONNECTION_URL, { useNewUrlParser: true, useUnifiedTopology: true })
  .then(() => app.listen(PORT, () => console.log(`Server Running on Port: http://localhost:${PORT}`)))
  .catch((error: Error) => console.log(`${error} did not connect`));

// Disable deprecated findAndModify function in Mongoose
mongoose.set('useFindAndModify', false);
