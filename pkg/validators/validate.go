package validators

import (
	"errors"
)

// AccessToken validates an AccessToken key.
func AccessToken(input string) error {
	if len(input) < 64 {
		return errors.New("Access Key key is too short, must be at least 64 characters long")
	}

	return nil
}
