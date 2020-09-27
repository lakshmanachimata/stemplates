package modals

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserName       string
	UserType       uint //  1 -admin ,2 - user 3 - contract , 4 -read access ,5 terminated (but data kept), 6 - terminated with data migrated (data migrated), 7 terminated  permanantly
	role           string
	UserMobile     string
	UserEmail      string
	UserCreated    time.Time
	UserStatus     string // active, inactive,
	lastLogin      time.Time
	userTerminated time.Time
	createdBy      primitive.ObjectID
}
