import mongoose, { Document, Schema, Model } from 'mongoose';

// Define an interface that describes the structure of a post document
interface IPost extends Document {
    title: string;
    message: string;
    creator: string;
    tags: string[];
    selectedFile: string;
    likeCount: number;
    createdAt: Date;
}

// Define the schema with type annotations
const postSchema: Schema<IPost> = new mongoose.Schema({
    title: { type: String, required: true },
    message: { type: String, required: true },
    creator: { type: String, required: true },
    tags: { type: [String], required: true },
    selectedFile: { type: String },
    likeCount: {
        type: Number,
        default: 0,
    },
    createdAt: {
        type: Date,
        default: new Date(),
    },
});

// Create the Post model
const PostMessage: Model<IPost> = mongoose.model<IPost>('PostMessage', postSchema);

export default PostMessage;
