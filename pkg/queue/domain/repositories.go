package domain

type QueueRepository interface {
	SaveQueue(slackId string, channelId string, queue *Queue)
	RestoreQueue(slackId string, channelId string) *Queue
}

type SlotRepository interface {
	SaveSlots(slackId string, channelId string, slots []*Slot)
	RestoreSlots(slackId string, channelId string) []*Slot
}
