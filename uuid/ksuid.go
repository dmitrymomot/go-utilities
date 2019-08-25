package uuid

import (
	"github.com/segmentio/ksuid"
)

type (
	// UUID structure
	UUID struct {
		uuid ksuid.KSUID
	}
)

// String returns generated uuid v1 string
func (u UUID) String() string {
	return u.uuid.String()
}

// New creates a new UUID object
func New() UUID {
	return UUID{ksuid.New()}
}
