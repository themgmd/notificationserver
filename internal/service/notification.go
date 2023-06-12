package service

import (
	"context"
	"notifier/internal/repository/dao"
)

type notificationRepo interface {
	Insert(ctx context.Context, notification dao.Notification) error
	GetUndelivered(ctx context.Context) ([]dao.Notification, error)
	GetAll(ctx context.Context) ([]dao.Notification, error)
	AddDeliveredTime(ctx context.Context, notificationId int) error
}

type NotificationService struct {
	repo notificationRepo
}

func NewNotificationService(repo notificationRepo) *NotificationService {
	return &NotificationService{repo: repo}
}
