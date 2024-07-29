package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//define structure/model to insert it in modngodb database
type Netflix struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
