package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSlotReturnErrorWhenSlotIsInvalid(t *testing.T) {
	slot, err := NewSlot("   ", nil)

	assert.Nil(t, slot)
	assert.NotNil(t, err)
}

func TestNewSlotReturnSlotWhenItsValid(t *testing.T) {
	slot, err := NewSlot("New Slot", nil)

	assert.Nil(t, err)
	assert.NotNil(t, slot)
}

func TestSimilarSlotsAreEqual(t *testing.T) {
	first, _ := NewSlot("a", nil)
	second, _ := NewSlot("a", nil)

	assert.True(t, first.Equals(second))
}
