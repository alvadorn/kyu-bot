package domain

import (
	"errors"
	"strings"
)

type Queue struct {
	data []string
}

func NewQueue(elements ...string) *Queue {
	queue := &Queue{}

	if elements != nil {
		queue.data = elements
	}

	return queue
}

func (queue *Queue) Enqueue(element string) error {
	trimmedElement := strings.TrimSpace(element)

	if trimmedElement == "" {
		return errors.New("Element can't be null")
	}

	queue.data = append(queue.data, trimmedElement)
	return nil
}

func (queue *Queue) Dequeue() string {
	if len(queue.data) > 0 {
		element := queue.data[0]
		queue.data = queue.data[1:]
		return element
	}
	return ""
}

func (queue *Queue) Size() int {
	return len(queue.data)
}

func (queue *Queue) Clear() {
	queue.data = []string{}
}

func (queue Queue) Elements() []string {
	dataCopy := make([]string, len(queue.data))
	copy(dataCopy, queue.data)

	return dataCopy
}
