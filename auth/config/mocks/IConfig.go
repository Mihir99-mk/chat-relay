// Code generated manually for testing purposes.
package mocks

import (
	"auth/config"
	"auth/ent/entgen"

	"github.com/stretchr/testify/mock"
	// Replace this with your actual DB type
	// if DB() returns *ent.Client
)

// IConfig is an autogenerated mock type for the config.IConfig interface
type IConfig struct {
	mock.Mock
}

// Env mocks the Env() method of config.IConfig
func (m *IConfig) Env() config.IEnv {
	args := m.Called()
	return args.Get(0).(config.IEnv)
}

// DB mocks the DB() method of config.IConfig
func (m *IConfig) DB() *entgen.Client {
	args := m.Called()
	return args.Get(0).(*entgen.Client)
}
