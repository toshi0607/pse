package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/toshi0607/pse/subcmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run() error {
	c, err := subcmd.Repository().Find(os.Args[1])
	if err != nil {
		log.Fatalf("failed to find cmd: %s err:%v", os.Args[1], err)
	}
	return c.Run(os.Args[1:])
}

func publish() error {
	projectID := os.Args[2]
	subcmd := os.Args[3]
	topicID := os.Args[4]

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	switch subcmd {
	// ① go run main.go pub testProject create testTopic
	case "create":
		topic, err := client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		fmt.Printf("Topic %v created.\n", topic)
	// ③ go run main.go pub testProject publish testTopic
	case "publish":
		topic := client.Topic(topicID)
		defer topic.Stop()
		for i := 0; i < 10; i++ {
			r := topic.Publish(ctx, &pubsub.Message{
				Data: []byte(fmt.Sprintf("hello world: %d", i)),
			})
			id, err := r.Get(ctx)
			if err != nil {
				log.Fatalf("Failed to get message id: %s", err)
			}
			fmt.Printf("message published: %s\n", id)
		}
	}

	return nil
}

func subscribe() error {
	projectID := os.Args[2]
	subcmd := os.Args[3]

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	switch subcmd {
	// ② go run main.go sub testProject create testTopic testSubscription
	case "create":
		topicID := os.Args[4]
		subscriptionID := os.Args[5]

		topic, err := client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		sub, err := client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
		}

		fmt.Printf("Subscription %v created.\n", sub)

	// ④ go run main.go sub testProject receive testSubscription
	case "receive":
		subscriptionID := os.Args[4]

		sub := client.Subscription(subscriptionID)
		err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			fmt.Printf("message received: %v\n", string(m.Data))
			m.Ack()
		})
		if err != context.Canceled {
			log.Fatalf("Failed to reveive message: %v", err)
		}
	}

	return nil
}
