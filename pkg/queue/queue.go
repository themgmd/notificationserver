package queue

import "github.com/onemgvv/notificationserver/internal/domain"

type Queue struct {
	notifications map[string]chan domain.Notification
}

func New() *Queue {
	return &Queue{
		notifications: make(map[string]chan domain.Notification, 10),
	}
}

func (q *Queue) Put(notification domain.Notification) {
	userId := notification.UserID.String()

	_, exists := q.notifications[userId]
	if !exists {
		q.notifications[userId] = make(chan domain.Notification, 10)
	}

	q.notifications[userId] <- notification
}

func (q *Queue) Get(userId string) <-chan domain.Notification {
	channel, exists := q.notifications[userId]
	if !exists {
		q.notifications[userId] = make(chan domain.Notification, 10)
	}

	return channel
}
