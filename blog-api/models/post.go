package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {	
	ID 	 primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	Tags	[]string           `bson:"tags"`
}