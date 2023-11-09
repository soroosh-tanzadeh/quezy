package queue

import "errors"

type Manager struct {
	queues       map[string]Queue
	defaultQueue string
}

func (q *Manager) RegisterQueue(name string, queue Queue) error {
	_, ok := q.queues[name]
	if ok {
		return errors.New("queue is already registered")
	}
	q.queues[name] = queue
	return nil
}

func (q *Manager) SetDefaultQueue(queueName string) error {
	_, ok := q.queues[queueName]
	if !ok {
		return errors.New("queue not registered")
	}
	q.defaultQueue = queueName
	return nil
}

func (q *Manager) GetDefaultQueue() Queue {
	return q.queues[q.defaultQueue]
}
