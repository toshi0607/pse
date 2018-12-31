package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

type Subscriber struct {
	client         *pubsub.Client
	projectID      string
	topicID        string
	subscriptionID string
}

func NewSubscriber(projectID, topicID, subscriptionID string) (*Subscriber, error) {
	ctx := context.Background()
	cli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		client:         cli,
		projectID:      projectID,
		topicID:        topicID,
		subscriptionID: subscriptionID,
	}, nil
}

func (s *Subscriber) CreateSubscription(ctx context.Context) (*pubsub.Subscription, error) {
	topic := s.client.Topic(s.topicID)
	sub, err := s.client.CreateSubscription(ctx, s.subscriptionID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *Subscriber) ReceiveSampleMessages(ctx context.Context) error {
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
