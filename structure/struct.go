package structure

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feedbackdata struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	LinkedIn string             `json:"linkedin,omitempty" bson:"linkedin,omitempty"`
	Company  string             `json:"company,omitempty" bson:"company,omitempty"`
}
