package tdd

import (
	"errors"
)

var ErrUsernameEmpty = errors.New("username is empty")
var ErrPasswordEmpty = errors.New("password is empty")

// Step 1: Create a new function called Validate() that returns a boolean value.
func (a *Account) Validate() error {
	// Step 4: Implement logic to validate the account.
	if a.Username == "" {
		return ErrUsernameEmpty
	}
	if a.Password == "" {
		return ErrPasswordEmpty
	}
	return nil
}
