package service

import (
	"context"

	"github.com/onemgvv/notificationserver/internal/domain"
)

type Notification interface {
	ProcessMessage(ctx context.Context, notification domain.Notification) error
}
