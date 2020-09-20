package flags

import "errors"

type EnvironmentFlag struct {
	env string
}

func (ef *EnvironmentFlag) String() string {
	// Default value is 'sandbox'
	if ef.env == "" {
		return "sandbox"
	}
	return ef.env
}

func (ef *EnvironmentFlag) Set(value string) error {
	switch value {
	case "sandbox", "production":
		ef.env = value
		return nil
	default:
		return errors.New("Invalid environment. Must be sandbox or production.")
	}
}

func (ef *EnvironmentFlag) Type() string {
	return "string"
}
