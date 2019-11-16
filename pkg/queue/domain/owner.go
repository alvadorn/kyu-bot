package domain

import (
	"errors"
	"strings"
)

type SlotOwner struct {
	owner string
}

func NewSlotOwner(ownerName string) (*SlotOwner, error) {
	trimmedOwnerName := strings.TrimSpace(ownerName)

	if len(trimmedOwnerName) == 0 {
		return nil, errors.New("SlotOwner can't be nullable")
	}

	return &SlotOwner{trimmedOwnerName}, nil
}

func (so SlotOwner) OwnerName() string {
	return so.owner
}
