package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSlotOwnerValid(t *testing.T) {
	owner, err := NewSlotOwner("NewOwner")

	assert.Nil(t, err)
	assert.NotNil(t, owner)
}

func TestNewSlotOwnerInvalid(t *testing.T) {
	owner, err := NewSlotOwner("    ")

	assert.Nil(t, owner)
	assert.NotNil(t, err)
}
