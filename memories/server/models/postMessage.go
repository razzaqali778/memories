package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Post represents the structure of a post document in MongoDB
type Post struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Message      string             `json:"message" bson:"message"`
	Creator      string             `json:"creator" bson:"creator"`
	Tags         []string           `json:"tags" bson:"tags"`
	SelectedFile string             `json:"selectedFile" bson:"selectedFile"`
	LikeCount    int                `json:"likeCount" bson:"likeCount"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
}
