package validators

import (
	"encoding/json"
	"errors"
)

// AccessToken validates an AccessToken key.
func AccessToken(input string) error {
	if len(input) < 64 {
		return errors.New("Access Key key is too short, must be at least 64 characters long")
	}

	return nil
}

// IsValidJSON check passed JSON string for validity
func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
