package models

import "gopkg.in/mgo.v2/bson"

type Console struct {
	Id           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string        `json:"name,omitempty" bson:"name,omitempty"`
	Manufacturer string        `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`
	Generation   int           `json:"generation,omitempty" bson:"generation,omitempty"`
	FormFactor   string        `json:"formfactor,omitempty" bson:"formfactor,omitempty"`
}
