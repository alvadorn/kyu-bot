package slot

import (
	"github.com/alvadorn/kyu_bot/pkg/queue/domain"
	"github.com/go-redis/redis/v7"
	"sort"
)

type SlotRepository struct {
	dbClient *redis.Client
}

func NewSlotRepository(client *redis.Client) *SlotRepository {
	return &SlotRepository{
		client,
	}
}

func (sr SlotRepository) SaveSlots(slackId, channelId string, slots []*domain.Slot) {
	for _, slot := range slots {
		key := keySlot(slackId, channelId, slot)

		sr.dbClient.HMSet(key, toDBFormat(slot))
	}
}

func (sr SlotRepository) RestoreSlots(slackId, channelId string) []*domain.Slot {
	key := keyGeneralSlot(slackId, channelId)

	keys, err := sr.dbClient.Keys(key).Result()

	if err != nil {
		return []*domain.Slot{}
	}

	var slots []*domain.Slot

	for _, keySlot := range keys {
		slotMap, _ := sr.dbClient.HGetAll(keySlot).Result()

		if slotMap != nil {
			slots = append(slots, toSlotFormat(slotMap))
		}
	}

	sort.Stable(domain.SlotAscending(slots))

	return slots
}

func keyGeneralSlot(slackId, channelId string) string {
	return slackId + ":" + channelId + ":slot:*"
}

func keySlot(slackId, channelId string, slot *domain.Slot) string {
	return slackId + ":" + channelId + ":slot:" + slot.ID()
}
