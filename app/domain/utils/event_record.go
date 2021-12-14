package utils

import "time"

const (
	CREATED_AT string = "createdAt"
	UPDATED_AT string = "updatedAt"
	DELETED_AT string = "deletedAt"
	ACTIVE     string = "active"
)

type EventRecord struct {
	CreatedAt *time.Time `bson:"created_At"`
	UpdatedAt *time.Time `bson:"updated_At"`
	DeletedAt *time.Time `bson:"deleted_At"`
	Active    bool       `bson:"active"`
}
