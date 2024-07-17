package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatementVersion struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	StatementID primitive.ObjectID `bson:"statement_id"`
	Date        time.Time          `bson:"date"`
	Signer      string             `bson:"signer"`
	FWD         string             `bson:"fwd"`
	FileLink    string             `bson:"file_link"`
}
