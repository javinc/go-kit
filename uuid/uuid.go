package uuid

import uuid "github.com/satori/go.uuid"

// New creates new v4 UUID.
func New() string {
	return uuid.NewV4().String()
}
