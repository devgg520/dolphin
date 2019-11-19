package model

import (
	"fmt"

	"gopkg.in/guregu/null.v3"

	"github.com/2637309949/dolphin/cli/platform/util"
	"golang.org/x/crypto/scrypt"
)

// SetPassword Method to set salt and hash the password for a user
func (u *User) SetPassword(password string) {
	b, err := util.RandomBytes(16)
	if err != nil {
		panic(err)
	}
	u.Salt = null.StringFrom(fmt.Sprintf("%x", b))
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	u.Password = null.StringFrom(fmt.Sprintf("%x", dk))
}

// ValidPassword Method to check the entered password is correct or not
func (u *User) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return u.Password.String == fmt.Sprintf("%x", dk)
}
