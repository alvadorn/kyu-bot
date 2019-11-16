package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlottedQueueAddNewSlot(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, false)

	slot, _ := NewSlot(uuid.New(), "A", nil)

	err := sq.AddNewSlot(slot)

	assert.Nil(t, err)
}

func TestSlottedQueueAddSlotRepeatablyWillFail(t *testing.T) {
	slot, _ := NewSlot(uuid.New(), "A", nil)
	slots := []*Slot{slot}
	sq := NewSlottedQueue(nil, slots, false)

	repeatedSlot, _ := NewSlot(uuid.New(), "A", nil)

	err := sq.AddNewSlot(repeatedSlot)

	assert.NotNil(t, err)
}

func TestSlottedQueueAddSlotOnSingleWillFail(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, true)

	repeatedSlot, _ := NewSlot(uuid.New(), "A", nil)

	err := sq.AddNewSlot(repeatedSlot)

	assert.NotNil(t, err)
}

func TestSlottedQueueRemoveSlotWillFailWhenNotFound(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, false)

	err := sq.RemoveSlotByID(uuid.New().String())

	assert.NotNil(t, err)
}

func TestSlottedQueueRemoveSlot(t *testing.T) {
	uuidv4 := uuid.New()
	slot, _ := NewSlot(uuidv4, "A", nil)
	slots := []*Slot{slot}
	sq := NewSlottedQueue(nil, slots, false)

	err := sq.RemoveSlotByID(uuidv4.String())

	assert.Nil(t, err)
}

func TestSlottedQueueRemoveSlotFailForSingle(t *testing.T) {
	uuidv4 := uuid.New()
	slot, _ := NewSlot(uuidv4, "A", nil)
	slots := []*Slot{slot}

	sq := NewSlottedQueue(nil, slots, true)

	err := sq.RemoveSlotByID(uuidv4.String())

	assert.NotNil(t, err)
}
