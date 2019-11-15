package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlotNewSlotReturnErrorWhenSlotIsInvalid(t *testing.T) {
	slot, err := NewSlot("   ", nil)

	assert.Nil(t, slot)
	assert.NotNil(t, err)
}

func TestSlotNewSlotReturnSlotWhenItsValid(t *testing.T) {
	slot, err := NewSlot("New Slot", nil)

	assert.Nil(t, err)
	assert.NotNil(t, slot)
}

func TestSlotSimilarSlotsAreEqual(t *testing.T) {
	first, _ := NewSlot("a", nil)
	second, _ := NewSlot("a", nil)

	assert.True(t, first.Equals(second))
}

func TestSlotReleaseOwner(t *testing.T) {
	owner, _ := NewSlotOwner("B")
	slot, _ := NewSlot("A", owner)

	assert.False(t, slot.IsEmpty())

	slot.ReleaseOwner()

	assert.True(t, slot.IsEmpty())
}

func TestSlotName(t *testing.T) {
	slot, _ := NewSlot("A", nil)

	assert.Equal(t, slot.Name(), "A")
}

func TestSlotOwnerChanges(t *testing.T) {
	slot, _ := NewSlot("A", nil)

	owner, _ := NewSlotOwner("B")

	assert.True(t, slot.IsEmpty())

	err := slot.NewOwner(owner)

	assert.Nil(t, err)
	assert.False(t, slot.IsEmpty())
}

func TestSlowOwnerDontChange(t *testing.T) {
	owner, _ := NewSlotOwner("B")
	slot, _ := NewSlot("A", owner)

	newOwner, _ := NewSlotOwner("C")

	assert.False(t, slot.IsEmpty())

	err := slot.NewOwner(newOwner)

	assert.NotNil(t, err)

}
