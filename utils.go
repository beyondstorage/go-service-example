package example

import (
	"github.com/aos-dev/go-storage/v3/services"
	"github.com/aos-dev/go-storage/v3/types"
)

// Storage is the example client.
type Storage struct {
	pairPolicy   types.PairPolicy
	defaultPairs DefaultStoragePairs
}

// String implements Storager.String
func (s *Storage) String() string {
	panic("implement me")
}

// NewStorager will create Storager only.
func NewStorager(pairs ...types.Pair) (types.Storager, error) {
	panic("implement me")
}

// formatError converts errors returned by SDK into errors defined in go-storage and go-service-*.
// The original error SHOULD NOT be wrapped.
func (s *Storage) formatError(op string, err error, path ...string) error {
	if _, ok := err.(services.AosError); ok {
		return err
	}

	panic("implement me")
}
