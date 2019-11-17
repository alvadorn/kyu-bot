package slot

import (
	"github.com/alvadorn/kyu_bot/pkg/queue/domain"
	"github.com/google/uuid"
)

const ID = "ID"
const owner = "owner"
const name = "name"

func toDBFormat(slot *domain.Slot) map[string]interface{} {
	return map[string]interface{}{
		ID:    slot.ID(),
		name:  slot.Name(),
		owner: slot.CurrentOwner(),
	}
}

func toSlotFormat(slotMap map[string]string) *domain.Slot {
	id := uuid.MustParse(slotMap[ID])

	owner, _ := domain.NewSlotOwner(slotMap[owner])

	slot, _ := domain.NewSlot(id, slotMap[name], owner)

	return slot
}
