package domain

type SlottedQueueRepository interface {
	recoverSlottedQueueInfo(slackId string, chatId string) (*SlottedQueue, error)
	saveSlottedQueue(slackId string, chatId string)
}

type QueueRepository interface {
	saveQueue(slackId string, channelId string, queue *Queue)
	restoreQueue(slackId string, channelId string) *Queue
}

type SlotRepository interface {
	saveSlots(slackId string, channelId string, slots []*Slot)
	restoreQueue(slackId string, channelId string) []*Slot
}
