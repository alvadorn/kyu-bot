package domain

import (
	"errors"
	"strings"
)

type Slot struct {
	name  string
	owner *SlotOwner
}

func NewSlot(name string, owner *SlotOwner) (*Slot, error) {
	trimmedName := strings.TrimSpace(name)

	if len(trimmedName) == 0 {
		return nil, errors.New("Slot name cant be null")
	}

	return &Slot{
		trimmedName,
		owner,
	}, nil
}

func (slot Slot) Equals(otherSlot *Slot) bool {
	return slot.name == otherSlot.Name()
}

func (slot Slot) Name() string {
	return slot.name
}

func (slot Slot) IsEmpty() bool {
	return slot.owner == nil
}

func (slot *Slot) ReleaseOwner() {
	slot.owner = nil
}

func (slot *Slot) NewOwner(owner *SlotOwner) error {
	if slot.owner == nil {
		slot.owner = owner
		return nil
	}

	return errors.New("Already has an owner, wait until the owner release its position")
}