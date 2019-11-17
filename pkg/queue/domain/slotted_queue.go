package domain

import (
	"errors"
	"github.com/google/uuid"
)

type SlottedQueueOptions struct {
	slots  []*Slot
	single bool
	queue  *Queue
}

type SlottedQueue struct {
	SlottedQueueOptions
	SlackID   string
	ChannelID string
}

func NewSlottedQueue(slackID, channelID string, options SlottedQueueOptions) *SlottedQueue {
	options.setDefault()

	slottedQueue := &SlottedQueue{
		options,
		slackID,
		channelID,
	}

	if len(slottedQueue.slots) == 0 {
		if slottedQueue.single {
			slot, _ := NewSlot(uuid.New(), singleSlotKey, nil)

			slottedQueue.slots = append(slottedQueue.slots, slot)
		}
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

func (options *SlottedQueueOptions) setDefault() {
	if options.slots == nil {
		options.slots = []*Slot{}
	}

	if options.queue == nil {
		options.queue = NewQueue()
	}
}
