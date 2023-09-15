package factory

import (
	"github.com/stretchr/testify/mock"
)

// MockFactory adalah implementasi mock dari Factory
type Factory struct {
	mock.Mock
}

// SetupDb adalah metode simulasi untuk SetupDb dalam Factory
func (m *Factory) SetupDb() {
	m.Called()
}

// SetupRepository adalah metode simulasi untuk SetupRepository dalam Factory
func (m *Factory) SetupRepository() {
	m.Called()
}

// NewFactory adalah metode simulasi untuk NewFactory dalam Factory
func (m *Factory) NewFactory() *Factory {
	args := m.Called()
	return args.Get(0).(*Factory)
}
