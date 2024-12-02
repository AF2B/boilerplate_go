package mocks

import (
	"errors"
)

type MockUserRepository struct {
	CalledWith string
}

func (m *MockUserRepository) SaveUser(name string) error {
	m.CalledWith = name
	if name == "error" {
		return errors.New("mock error")
	}
	return nil
}
