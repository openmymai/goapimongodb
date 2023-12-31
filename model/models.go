package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie" bson:"movie"`
	Watched bool               `json:"watched" bson:"watched"`
}
