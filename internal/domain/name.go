package domain

import "errors"

type Name string

func (n Name) String() string {
	return string(n)
}

var EmptyName = errors.New("name is empty")

func ParseName(value string) (Name, error) {
	name := Name(value)

	if err := name.validate(); err != nil {
		return "", err
	}

	return name, nil
}

func (n Name) validate() error {
	if len(n) == 0 {
		return EmptyName
	}

	return nil
}
