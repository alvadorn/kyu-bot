package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeueReturnsNilIfQueueIsEmpty(t *testing.T) {
	q := NewQueue()
	value := q.Dequeue()

	assert.Empty(t, value)
}

func TestQueueWorksLikeFIFO(t *testing.T) {
	q := NewQueue()
	q.Enqueue("1")
	q.Enqueue("2")

	assert.Equal(t, q.Dequeue(), "1")
	assert.Equal(t, q.Dequeue(), "2")
}

func TestQueueSize(t *testing.T) {
	q := NewQueue()

	assert.Zero(t, q.Size())

	q.Enqueue("1")

	assert.Equal(t, q.Size(), 1)
}

func TestQueueClear(t *testing.T) {
	q := NewQueue()

	q.Enqueue("1")
	assert.Equal(t, q.Size(), 1)

	q.Clear()
	assert.Zero(t, q.Size())
}

func TestElements(t *testing.T) {
	q := NewQueue("1")

	assert.Equal(t, len(q.Elements()), 1)
}
