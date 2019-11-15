package domain

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Enqueue(element interface{}) {
	queue.data = append(queue.data, element)
}

func (queue *Queue) Dequeue() interface{} {
	if (len(queue.data) > 0) {
		element := queue.data[0]
		queue.data = queue.data[1:]
		return element
	}
	return nil
}

func (queue *Queue) Size() int {
	return len(queue.data)
}

func (queue *Queue) Clear() {
	queue.data = []interface{}{}
}