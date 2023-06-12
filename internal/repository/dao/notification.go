package dao

import (
	"time"

	"github.com/onemgvv/notificationserver/internal/domain"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

const NotificationTableName = "notifications"

var (
	InsertNotificationColumns = []string{
		"id",
		"user_id",
		"status_id",
		"header",
		"body",
	}

	SelectNotificationColumns = []string{
		"id",
		"user_id",
		"status_id",
		"header",
		"body",
		"delivered_at",
		"created_at",
	}
)

type Notification struct {
	ID          int       `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	StatusID    int       `db:"status_id"`
	Header      string    `db:"header"`
	Body        string    `db:"body"`
	DeliveredAt null.Time `db:"delivered_at"`
	CreatedAt   time.Time `db:"created_at"`
}

func (n Notification) TableName() string {
	return NotificationTableName
}

func (n Notification) ToDomain() *domain.Notification {
	return &domain.Notification{
		UserID: n.UserID,
		Status: domain.Status(n.StatusID),
		Header: n.Header,
		Body:   n.Body,
	}
}

func (n Notification) GetInsertData() []interface{} {
	return []interface{}{
		n.ID, n.UserID, n.StatusID,
		n.Header, n.Body,
	}
}

func NotificationFromDomain(message domain.Notification) *Notification {
	return &Notification{
		UserID:   message.UserID,
		StatusID: int(message.Status),
		Header:   message.Header,
		Body:     message.Body,
	}
}
