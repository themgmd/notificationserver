package repository

import (
	"context"
	"fmt"
	"notifier/internal/repository/dao"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (nr NotificationRepository) Insert(ctx context.Context, notification dao.Notification) error {
	sql, args, err := squirrel.Insert(dao.NotificationTableName).
		Columns(dao.InsertNotificationColumns...).
		Values(notification.GetInsertData()...).
		ToSql()
	if err != nil {
		return fmt.Errorf("squirrel.Insert: %w", err)
	}

	_, err = nr.db.ExecContext(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("nr.db.ExecContext: %w", err)
	}

	return nil
}

func (nr NotificationRepository) GetUndelivered(ctx context.Context) ([]dao.Notification, error) {
	var notifications []dao.Notification

	sql, args, err := squirrel.
		Select(dao.SelectNotificationColumns...).
		From(dao.NotificationTableName).
		Where(squirrel.NotEq{"delivared_at": nil}).
		ToSql()
	if err != nil {
		return notifications, fmt.Errorf("squirrel.SelectAll: %w", err)
	}

	err = nr.db.SelectContext(ctx, &notifications, sql, args)
	if err != nil {
		return notifications, fmt.Errorf("nr.db.SelectContext: %w", err)
	}

	return notifications, nil
}

func (nr NotificationRepository) GetAll(ctx context.Context) ([]dao.Notification, error) {
	var notifications []dao.Notification

	sql, args, err := squirrel.
		Select(dao.SelectNotificationColumns...).
		From(dao.NotificationTableName).
		ToSql()
	if err != nil {
		return notifications, fmt.Errorf("squirrel.SelectAll: %w", err)
	}

	err = nr.db.SelectContext(ctx, &notifications, sql, args)
	if err != nil {
		return notifications, fmt.Errorf("nr.db.SelectContext: %w", err)
	}

	return notifications, nil
}

func (nr NotificationRepository) AddDeliveredTime(ctx context.Context, notificationId int) error {
	sql, args, err := squirrel.
		Update(dao.NotificationTableName).
		Set("delivered_at", time.Now()).
		Where(squirrel.Eq{"id": notificationId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("squirrel.Update: %w", err)
	}

	_, err = nr.db.ExecContext(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("nr.db.ExecContext: %w", err)
	}

	return nil
}
