package model

type Log struct {
	Id      string `bson:"_id" json:"_id"`
	Message string `bson:"message" json:"message"`
}
