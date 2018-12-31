package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

type Subscriber interface {
	CreateSubscription(ctx context.Context, subscriptionID string, topicID string) (*pubsub.Subscription, error)
	ReceiveSampleMessages(ctx context.Context, subscriptionID string) error
	Init(ctx context.Context, projectID string) error
}

type subscriber struct {
	client *pubsub.Client
}

func NewSubscriber() Subscriber {
	return &subscriber{}
}

func (s *subscriber) CreateSubscription(ctx context.Context, subscriptionID, topicID string) (*pubsub.Subscription, error) {
	topic := s.client.Topic(topicID)
	sub, err := s.client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *subscriber) ReceiveSampleMessages(ctx context.Context, subscriptionID string) error {
	sub := s.client.Subscription(subscriptionID)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("message received: %v\n", string(m.Data))
		m.Ack()
	})
	if err != context.Canceled {
		return err
	}

	return nil
}

func (s *subscriber) Init(ctx context.Context, projectID string) error {
	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	s.client = c

	return nil
}
