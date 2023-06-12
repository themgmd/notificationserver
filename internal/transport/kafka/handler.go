package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/onemgvv/notificationserver/internal/domain"
)

func (ks KafkaServer) ListenNotifications(ctx context.Context) {
	for {
		msg, err := ks.reader.FetchMessage(ctx)
		if err != nil {
			err = fmt.Errorf("ks.reader.FetchMessage: %w", err)
			log.Printf("[ERROR]: %s", err.Error())
			break
		}

		key := string(msg.Key)
		if !strings.Contains(key, "notification/") {
			continue
		}

		var notification domain.Notification
		err = json.Unmarshal(msg.Value, &notification)
		if err != nil {
			err = fmt.Errorf("json.Unmarshal: %w", err)
			log.Printf("[ERROR]: %s", err.Error())
			continue
		}

		ks.notificationService.ProcessMessage(ctx, notification)

		err = ks.reader.CommitMessages(ctx, msg)
		if err != nil {
			err = fmt.Errorf("ks.reader.CommitMessages: %w", err)
			log.Printf("[ERROR]: %s", err.Error())
			continue
		}
	}

	err := ks.reader.Close()
	if err != nil {
		err = fmt.Errorf("ks.reader.Close: %w", err)
		log.Printf("[ERROR]: %s", err.Error())
	}
}
