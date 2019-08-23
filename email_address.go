package valueobject

import (
	"errors"
	"regexp"

)

// EmailAddress errors
var (
	ErrInvalidEmailAddress = errors.New("not a valid email address")
)

// EmailAddress represents a valid email address
type EmailAddress struct {
	value string
}

// NewEmailAddress creates a new email address
func NewEmailAddress(email string) (EmailAddress, error) {
	var n EmailAddress
	match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)
	if !match {
		return n, ErrInvalidEmailAddress
	}
	n.value = email

	return n, nil
}

// String returns string representation of the email address
func (e EmailAddress) String() string {
	return e.value
}

// Equals checks that two email addresses are the same
func (e EmailAddress) Equals(value Value) bool {
	otherEmailAddress, ok := value.(EmailAddress)
	return ok && e.value == otherEmailAddress.value
}