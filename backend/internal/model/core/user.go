package core

import (
	"encoding/base64"
	"fmt"

	"junction-brella/internal/constants"
)

type UserID string

// UserPassword describes user's password data
type UserPassword struct {
	Hash string `bson:"hash"`
	Salt string `bson:"salt"`
}

// User describes a user entity
type User struct {
	ID       UserID       `bson:"_id"`
	Email    string       `bson:"email"`
	Username string       `bson:"username"`
	Password UserPassword `bson:"password"`

	CreatedAt int64 `bson:"created_at"`           // unix timestamp
	UpdatedAt int64 `bson:"updated_at,omitempty"` // unix timestamp

	Mode        string `bson:"mode"`
	Occupation  string `bson:"occupation"`
	PlaceOfWork string `bson:"place_of_work"`
	PhotoURL    string `bson:"photo_url"`
}

// Init generates salt and hash with given password and fills corresponding fields
func (up *UserPassword) Init(password string) error {
	salt, err := getSalt()
	if err != nil {
		return fmt.Errorf("error generating salt: %s", err)
	}

	hash, err := getHash512(password, salt)
	if err != nil {
		return fmt.Errorf("error generating hash: %s", err)
	}

	up.Salt = base64.URLEncoding.EncodeToString(salt)
	up.Hash = base64.URLEncoding.EncodeToString(hash)

	return nil
}

// Validate checks if the given password is the one that is stored
func (up *UserPassword) Validate(password string) error {
	salt, err := base64.URLEncoding.DecodeString(up.Salt)
	if err != nil {
		return fmt.Errorf("error decoding user salt: %s", err)
	}

	hash, err := getHash512(password, salt)
	if err != nil {
		return fmt.Errorf("error generating hash: %s", err)
	}

	if base64.URLEncoding.EncodeToString(hash) != up.Hash {
		return constants.ErrPasswordMismatch
	}

	return nil
}
