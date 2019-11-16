package domain

import (
	"errors"
	"github.com/google/uuid"
)

type SlottedQueue struct {
	slots  []*Slot
	queue  *Queue
	single bool
}

func NewSlottedQueue(queue *Queue, slots []*Slot, singleSlot bool) *SlottedQueue {
	slottedQueue := &SlottedQueue{
		slots,
		queue,
		singleSlot,
	}

	if slots == nil {
		slottedQueue.slots = []*Slot{}
	}

	if len(slots) == 0 {
		if singleSlot {
			slot, _ := NewSlot(uuid.New(), singleSlotKey, nil)

			slottedQueue.slots = append(slottedQueue.slots, slot)
		}
	}

	if queue == nil {
		slottedQueue.queue = NewQueue()
	}

	return slottedQueue
}

func (slottedQueue *SlottedQueue) AddNewSlot(newSlot *Slot) error {
	if slottedQueue.single {
		return errors.New("It is not possible to add a slottedQueue on a single slottedQueue queue")
	}

	if !slottedQueue.isSlotNamedUnique(newSlot) {
		return errors.New("Already exists a slot with this name")
	}

	slottedQueue.slots = append(slottedQueue.slots, newSlot)

	return nil
}

func (sq *SlottedQueue) isSlotNamedUnique(comparingSlot *Slot) bool {
	unique := true

	for _, slot := range sq.slots {
		unique = unique && slot.Name() != comparingSlot.Name()
	}

	return unique
}

func (sq *SlottedQueue) RemoveSlotByID(id string) error {
	if sq.single {
		return errors.New("Can't remove slot for a single slot queue")
	}

	indexFound := -1

	for index, slot := range sq.slots {
		if slot.ID() == id {
			indexFound = index
		}
	}

	if indexFound != -1 {
		sq.slots = append(sq.slots[:indexFound], sq.slots[indexFound+1:]...)
		return nil
	}

	return errors.New("Slot not found")
}
