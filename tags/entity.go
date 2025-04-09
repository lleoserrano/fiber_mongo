package tags

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Tasks []string           `json:"-" bson:"tasks,omitempty"`
}
