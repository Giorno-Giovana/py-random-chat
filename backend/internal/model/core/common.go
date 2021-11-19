package core

import (
	"crypto/rand"
	"crypto/sha512"
)

const saltSize = 32

func getSalt() ([]byte, error) {
	passwordSalt := make([]byte, saltSize)
	_, err := rand.Read(passwordSalt)
	return passwordSalt, err
}

func getHash512(password string, salt []byte) ([]byte, error) {
	var passwordHash []byte
	sha512Hasher := sha512.New()
	if _, err := sha512Hasher.Write(append([]byte(password), salt...)); err != nil {
		return passwordHash, err
	}
	passwordHash = sha512Hasher.Sum(nil)
	return passwordHash, nil
}

// Status implements bitfield operations
type Status uint32

func (s *Status) Set(status Status) {
	*s |= status
}

func (s *Status) Has(status Status) bool {
	return *s&status == status
}

func (s *Status) Clear(status Status) {
	*s &^= status // bitclear operator
}
