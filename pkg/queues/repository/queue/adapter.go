package queue

import (
	"github.com/alvadorn/kyu_bot/pkg/queues/domain"
	"strings"
)

const separator = ","

func toDBFormat(queue *domain.Queue) string {
	elements := queue.Elements()
	return strings.Join(elements, separator)
}

func toQueueFormat(dbData string) *domain.Queue {
	elements := strings.Split(dbData, separator)

	if len(elements) == 1 && elements[0] == "" {
		elements = []string{}
	}

	return domain.NewQueue(elements...)
}
