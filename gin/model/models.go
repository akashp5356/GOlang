package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Mid       string             `json:"mid"`
	Moviename string             `json:"movie,omitempty"`
	Watched   bool               `json:"watched"`
}
