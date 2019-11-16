package domain

type SlottedQueueRepository interface {
	recoverSlottedQueueInfo(slackId string, chatId string) (*SlottedQueue, error)
	saveSlottedQueue(slackId string, chatId string)
}

type QueueRepository interface {
	SaveQueue(slackId string, channelId string, queue *Queue)
	RestoreQueue(slackId string, channelId string) *Queue
}

type SlotRepository interface {
	saveSlots(slackId string, channelId string, slots []*Slot)
	restoreQueue(slackId string, channelId string) []*Slot
}
