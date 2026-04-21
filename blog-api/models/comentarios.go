package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comentario struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	PostID  primitive.ObjectID `bson:"post_id"`
	Texto   string             `bson:"texto"`
}