package slot

import (
	"github.com/alvadorn/kyu_bot/pkg/queue/domain"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RepositorySuite struct {
	suite.Suite
	client *redis.Client

	sr *SlotRepository
}

func (suite *RepositorySuite) SetupSuite() {
	suite.client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	suite.client.FlushAll()
}

func (suite *RepositorySuite) SetupTest() {
	suite.sr = NewSlotRepository(suite.client)
}

func (suite *RepositorySuite) TearDownTest() {
	suite.client.FlushAll()
}

func (suite *RepositorySuite) TestSaveSlots() {
	t := suite.T()
	sr := suite.sr

	slot1, _ := domain.NewSlot(uuid.New(), "A", nil)
	slot2, _ := domain.NewSlot(uuid.New(), "B", nil)

	slots := []*domain.Slot{slot1, slot2}

	sr.SaveSlots("slackId", "channelId", slots)

	keyPattern := keyGeneralSlot("slackId", "channelId")
	savedSlots, _ := suite.client.Keys(keyPattern).Result()

	assert.Equal(t, len(savedSlots), 2)
}

func (suite *RepositorySuite) TestRestoreSlotsNotSaved() {
	t := suite.T()
	sr := suite.sr

	slots := sr.RestoreSlots("slackId", "channelId")

	assert.Empty(t, slots)
}

func (suite *RepositorySuite) TestRestorePreviousSavedSlots() {
	t := suite.T()
	sr := suite.sr

	uuidv4 := uuid.New()
	slot, _ := domain.NewSlot(uuidv4, "A", nil)
	key := keySlot("slackId", "channelId", slot)

	sr.dbClient.HMSet(key, map[string]interface{}{
		ID:    uuidv4.String(),
		name:  slot.Name(),
		owner: "",
	})

	slots := sr.RestoreSlots("slackId", "channelId")

	assert.NotEmpty(t, slots)
	recoveredSlot := slots[0]
	assert.Equal(t, recoveredSlot.Name(), slot.Name())
	assert.Equal(t, recoveredSlot.ID(), slot.ID())
	assert.Equal(t, recoveredSlot.CurrentOwner(), slot.CurrentOwner())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
