package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeueReturnsNilIfQueueIsEmpty(t *testing.T) {
	q := &Queue{}
	value := q.Dequeue()

	assert.Nil(t, value)
}

func TestQueueWorksLikeFIFO(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
}

func TestQueueSize(t *testing.T) {
	q := &Queue{}

	assert.Zero(t, q.Size())

	q.Enqueue(1)

	assert.Equal(t, q.Size(), 1)
}

func TestQueueClear(t *testing.T) {
	q := &Queue{}

	q.Enqueue(1)
	assert.Equal(t, q.Size(), 1)

	q.Clear()
	assert.Zero(t, q.Size())
}
