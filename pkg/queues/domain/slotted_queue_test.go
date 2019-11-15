package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlottedQueueAddNewSlot(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, false)

	slot, _ := NewSlot("A", nil)

	err := sq.AddNewSlot(slot)

	assert.Nil(t, err)
}

func TestSlottedQueueAddSlotRepeateadlyWillFail(t *testing.T) {
	slot, _ := NewSlot("A", nil)
	slots := []*Slot{slot}
	sq := NewSlottedQueue(nil, slots, false)

	repeatedSlot, _ := NewSlot("A", nil)

	err := sq.AddNewSlot(repeatedSlot)

	assert.NotNil(t, err)
}

func TestSlottedQueueAddSlotOnSingleWillFail(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, true)

	repeatedSlot, _ := NewSlot("A", nil)

	err := sq.AddNewSlot(repeatedSlot)

	assert.NotNil(t, err)
}

func TestSlottedQueueRemoveSlotWillFailWhenNotFound(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, false)

	err := sq.RemoveSlotByName("A")

	assert.NotNil(t, err)
}

func TestSlottedQueueRemoveSlot(t *testing.T) {
	slot, _ := NewSlot("A", nil)
	slots := []*Slot{slot}
	sq := NewSlottedQueue(nil, slots, false)

	err := sq.RemoveSlotByName("A")

	assert.Nil(t, err)
}

func TestSlottedQueueRemoveSlotFailForSingle(t *testing.T) {
	sq := NewSlottedQueue(nil, nil, true)

	err := sq.RemoveSlotByName("A")

	assert.NotNil(t, err)
}