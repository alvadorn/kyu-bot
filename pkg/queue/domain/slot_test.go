package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlotNewSlotReturnErrorWhenSlotIsInvalid(t *testing.T) {
	slot, err := NewSlot(uuid.New(), "   ", nil)

	assert.Nil(t, slot)
	assert.NotNil(t, err)
}

func TestSlotNewSlotReturnSlotWhenItsValid(t *testing.T) {
	slot, err := NewSlot(uuid.New(), "New Slot", nil)

	assert.Nil(t, err)
	assert.NotNil(t, slot)
}

func TestSlotSimilarSlotsAreEqual(t *testing.T) {
	uuidv4 := uuid.New()
	first, _ := NewSlot(uuidv4, "a", nil)
	second, _ := NewSlot(uuidv4, "a", nil)

	assert.True(t, first.Equals(second))
}

func TestSlotReleaseOwner(t *testing.T) {
	owner, _ := NewSlotOwner("B")
	slot, _ := NewSlot(uuid.New(), "A", owner)

	assert.False(t, slot.IsEmpty())

	slot.ReleaseOwner()

	assert.True(t, slot.IsEmpty())
}

func TestSlotName(t *testing.T) {
	slot, _ := NewSlot(uuid.New(), "A", nil)

	assert.Equal(t, slot.Name(), "A")
}

func TestSlotOwnerChanges(t *testing.T) {
	slot, _ := NewSlot(uuid.New(), "A", nil)

	owner, _ := NewSlotOwner("B")

	assert.True(t, slot.IsEmpty())

	err := slot.NewOwner(owner)

	assert.Nil(t, err)
	assert.False(t, slot.IsEmpty())
}

func TestSlowOwnerDontChange(t *testing.T) {
	owner, _ := NewSlotOwner("B")
	slot, _ := NewSlot(uuid.New(), "A", owner)

	newOwner, _ := NewSlotOwner("C")

	assert.False(t, slot.IsEmpty())

	err := slot.NewOwner(newOwner)

	assert.NotNil(t, err)
}
