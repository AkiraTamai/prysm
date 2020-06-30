package direct

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "keymanager-v2")

// DirectConfig --
type DirectConfig struct{}

// Keymanager implementation for direct keystores.
type Keymanager struct{}

// DefaultConfig for a direct keymanager implementation.
func DefaultConfig() *DirectConfig {
	return &DirectConfig{}
}

// NewKeymanager --
func NewKeymanager(ctx context.Context, cfg *DirectConfig) *Keymanager {
	return &Keymanager{}
}

// CreateAccount for a direct keymanager implementation.
func (dr *Keymanager) CreateAccount(ctx context.Context, password string) error {
	return errors.New("unimplemented")
}

// ConfigFile --
func (dr *Keymanager) ConfigFile(ctx context.Context) ([]byte, error) {
	return nil, errors.New("unimplemented")
}
