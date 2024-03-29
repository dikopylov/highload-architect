package users

import (
	"github.com/dikopylov/highload-architect/internal/model/types"
	"time"
)

type User struct {
	ID        types.UserID `db:"id"`
	Birthdate *time.Time   `db:"birthdate"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Biography string       `db:"biography"`
	City      string       `db:"city"`
	Password  string       `db:"password"`
	Age       uint         `db:"age"`
}

type Users []*User
