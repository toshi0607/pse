package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type Publisher interface {
	CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	DeleteTopic(ctx context.Context, topicID string) error
	PublishSampleMessage(ctx context.Context, topicID string) error
	Init(ctx context.Context, projectID string) error
}

type publisher struct {
	client *pubsub.Client
}

func NewPublisher() Publisher {
	return &publisher{}
}

func (p *publisher) CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	topic, err := p.client.CreateTopic(ctx, topicID)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (p *publisher) DeleteTopic(ctx context.Context, topicID string) error {
	topic := p.client.Topic(topicID)
	err := topic.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *publisher) PublishSampleMessage(ctx context.Context, topicID string) error {
	topic := p.client.Topic(topicID)
	defer topic.Stop()
	for i := 0; i < 10; i++ {
		r := topic.Publish(ctx, &pubsub.Message{
			Data: []byte(fmt.Sprintf("hello world: %d", i)),
		})
		id, err := r.Get(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("message published: %s\n", id)
	}

	return nil
}

func (s *publisher) Init(ctx context.Context, projectID string) error {
	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	s.client = c

	return nil
}
