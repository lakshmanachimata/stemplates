package modals

import "time"

type User struct {
	UserName    string
	UserType    string
	UserMobile  string
	UserEmail   string
	UserCreated time.Time
	UserStatus  string
}
