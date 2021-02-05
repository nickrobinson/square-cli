package config

import "errors"

type Environment string

const (
	Sandbox    Environment = "sandbox"
	Production             = "production"
)

func (e *Environment) String() string {
	switch *e {
	case Sandbox:
		return "sandbox"
	case Production:
		return "production"
	}
	return ""
}

func (e *Environment) Set(name string) error {
	if name == "sandbox" {
		*e = Sandbox
	} else if name == "production" {
		*e = Production
	} else {
		return errors.New("Invalid Environment Name")
	}
	return nil
}

func (e *Environment) Type() string {
	return "Environment"
}
