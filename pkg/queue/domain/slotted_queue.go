package domain

import "errors"

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
			slot, _ := NewSlot(singleSlotKey, nil)

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

	if !slottedQueue.isSlotUnique(newSlot) {
		return errors.New("Already exists a slot with this name")
	}

	slottedQueue.slots = append(slottedQueue.slots, newSlot)

	return nil
}

func (sq *SlottedQueue) isSlotUnique(comparingSlot *Slot) bool {
	unique := true

	for _, slot := range sq.slots {
		unique = unique && !slot.Equals(comparingSlot)
	}

	return unique
}

func (sq *SlottedQueue) RemoveSlotByName(name string) error {
	if sq.single {
		return errors.New("Can't remove slot for a single slot queue")
	}

	slotToBeFound, err := NewSlot(name, nil)

	indexFound := -1

	if err == nil {
		for index, slot := range sq.slots {
			if slot.Equals(slotToBeFound) {
				indexFound = index
			}
		}
	}

	if indexFound != -1 {
		sq.slots = append(sq.slots[:indexFound], sq.slots[indexFound+1:]...)
		return nil
	}

	return errors.New("Slot not found")
}
