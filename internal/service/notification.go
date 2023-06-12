package service

import (
	"context"
	"fmt"

	"github.com/onemgvv/notificationserver/internal/domain"
	"github.com/onemgvv/notificationserver/internal/repository/dao"
	"github.com/onemgvv/notificationserver/pkg/queue"
)

type notificationRepo interface {
	Insert(ctx context.Context, notification dao.Notification) error
	GetUndelivered(ctx context.Context) ([]dao.Notification, error)
	GetAll(ctx context.Context) ([]dao.Notification, error)
	AddDeliveredTime(ctx context.Context, notificationId int) error
}

type NotificationService struct {
	queue *queue.Queue
	repo  notificationRepo
}

func NewNotificationService(queue *queue.Queue, repo notificationRepo) *NotificationService {
	return &NotificationService{
		queue: queue,
		repo:  repo,
	}
}

func (ns NotificationService) ProcessMessage(ctx context.Context, notification domain.Notification) error {
	ns.queue.Put(notification)

	err := ns.repo.Insert(ctx, *dao.NotificationFromDomain(notification))
	if err != nil {
		return fmt.Errorf("ns.repo.Insert: %w", err)
	}

	return nil
}
