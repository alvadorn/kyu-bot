package domain

import (
	"errors"
	"strings"
)

type Slot struct {
	name string
	placement interface{}
}

func NewSlot(name string, placement interface{}) (*Slot, error) {
	trimmedName := strings.TrimSpace(name)

	if len(trimmedName) == 0 {
		return nil, errors.New("Slot name cant be null")
	}

	return &Slot{
		trimmedName,
		placement,
	}, nil
}

func (slot *Slot) Equals(otherSlot *Slot) bool {
	return slot.name == otherSlot.Name()
}

func (slot *Slot) Name() string {
	return slot.name
}