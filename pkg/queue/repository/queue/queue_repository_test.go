package queue

import (
	"github.com/alvadorn/kyu_bot/pkg/queue/domain"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RepositorySuite struct {
	suite.Suite
	client *redis.Client
	qr     *QueueRepository
}

func (suite *RepositorySuite) SetupSuite() {
	suite.client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	suite.client.FlushAll()
}

func (suite *RepositorySuite) SetupTest() {
	suite.qr = NewQueueRepository(suite.client)
}

func (suite *RepositorySuite) TearDownTest() {
	suite.client.FlushAll()
}

func (suite *RepositorySuite) TestNewQueueRepository() {
	t := suite.T()
	qr := NewQueueRepository(suite.client)

	assert.NotNil(t, qr)
	assert.IsType(t, &QueueRepository{}, qr)
}

func (suite *RepositorySuite) TestSaveQueue() {
	t := suite.T()
	qr := suite.qr
	client := suite.client

	queue := domain.NewQueue("1", "2", "3")

	qr.SaveQueue("slackId", "channelId", queue)

	key := queueKey("slackId", "channelId")
	assert.Equal(t, client.Exists(key).Val(), int64(1))
}

func (suite *RepositorySuite) TestRestoreQueueEmpty() {
	t := suite.T()
	qr := suite.qr

	queue := qr.RestoreQueue("slackId", "channelId")

	assert.Empty(t, queue.Elements())
	assert.Zero(t, queue.Size())
}

func (suite *RepositorySuite) TestRestoreAFilledEntry() {
	t := suite.T()
	qr := suite.qr
	client := suite.client

	key := queueKey("slackId", "channelId")
	client.Set(key, "5,6", 0).Result()

	queue := qr.RestoreQueue("slackId", "channelId")

	assert.NotEmpty(t, queue.Elements())
	assert.Equal(t, queue.Size(), 2)
	assert.ElementsMatch(t, queue.Elements(), []string{"5", "6"})
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
