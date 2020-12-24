package config

import "errors"

type Environment struct {
	name string
}

func (e *Environment) String() string {
	return e.name
}

func (e *Environment) Set(name string) error {
	if name == "production" || name == "sandbox" {
		e.name = name
		return nil
	} else {
		return errors.New("Invalid environment name, must be sandbox or production")
	}
}

func (e *Environment) Type() string {
	return "Environment"
}
