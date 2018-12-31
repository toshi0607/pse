package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type Publisher struct {
	client    *pubsub.Client
	projectID string
	topicID   string
}

func NewPublisher(projectID, topicID string) (*Publisher, error) {
	ctx := context.Background()
	cli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		client:    cli,
		projectID: projectID,
		topicID:   topicID,
	}, nil
}

func (p *Publisher) CreateTopic(ctx context.Context) (*pubsub.Topic, error) {
	topic, err := p.client.CreateTopic(ctx, p.topicID)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (p *Publisher) PublishSampleMessage(ctx context.Context) error {
	topic := p.client.Topic(p.topicID)
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
