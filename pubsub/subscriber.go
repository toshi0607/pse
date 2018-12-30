package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

type Subscriber struct {
	client         pubsub.Client
	projectID      string
	topicID        string
	subscriptionID string
}

func (s *Subscriber) CreateSubscription(ctx context.Context) (*pubsub.Subscription, error) {

	topic, err := s.client.CreateTopic(ctx, s.topicID)
	if err != nil {
		return nil, err
	}

	sub, err := s.client.CreateSubscription(ctx, s.subscriptionID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *Subscriber) ReceiveSampleMesssages(ctx context.Context) error {
	sub := s.client.Subscription(s.subscriptionID)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("message received: %v\n", string(m.Data))
		m.Ack()
	})
	if err != context.Canceled {
		return err
	}

	return nil
}
