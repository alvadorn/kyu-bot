package queue

import (
	"github.com/alvadorn/kyu_bot/pkg/queue/domain"
	"github.com/go-redis/redis/v7"
)

type QueueRepository struct {
	dbClient *redis.Client
}

func NewQueueRepository(client *redis.Client) *QueueRepository {
	return &QueueRepository{
		client,
	}
}

func (qr QueueRepository) SaveQueue(slackId, channelId string, queue *domain.Queue) {
	key := queueKey(slackId, channelId)

	qr.dbClient.Set(key, toDBFormat(queue), 0).Result()
}

func (qr QueueRepository) RestoreQueue(slackId, channelId string) *domain.Queue {
	key := queueKey(slackId, channelId)

	queueData, err := qr.dbClient.Get(key).Result()

	if err == nil {
		return toQueueFormat(queueData)
	}

	return toQueueFormat("")
}

func queueKey(slackId, channelId string) string {
	return slackId + ":" + channelId + ":" + "queue_elements"
}
