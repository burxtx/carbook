package users

import (
	"time"

	"github.com/burxtx/car/users/utils"
)

type UserID string

// User is the central of user domain model
// play as aggragate root
type User struct {
	ID         UserID
	Profile    Profile
	account    Account
	IsActive   bool
	LastLogin  time.Time
	DateJoined time.Time
	IsLocked   bool
}

// Profile is value object
type Profile struct {
	FirstName string
	LastName  string
	NickName  string
	Sign      string
	Residence string
	Sex       string
	Avatar    string
}

// Account is entity
type Account struct {
	Username string
	Phone    string
	Email    string
	Password Password
}

// Account is value object
type Password struct {
	passwdHash string
	passwdSalt string
}

// NewUser creates a new user
func NewUser(account Account) *User {
	return &User{
		ID:      UserID(utils.NewUserID()),
		account: account,
	}
}

// NewAccount creates a new account
func NewAccount(username string, password Password) Account {
	return Account{
		Username: username,
		Password: password,
	}
}

// Repository provides access to a user store.
type UserRepository interface {
	Create(user *User) error
	Find(userID string) (*User, error)
	List() []*User
}
