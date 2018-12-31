package subcmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/pkg/errors"
	tf "github.com/toshi0607/pse/flag"
	tps "github.com/toshi0607/pse/pubsub"
)

type (
	Sub struct {
		flagSet *flag.FlagSet
		opts    SubConfig
	}

	SubConfig struct {
		ProjectID, TopicID, SubscriptionID string
	}
)

func NewSub() *Sub {
	s := &Sub{}
	s.flagSet = tf.New(s.Name(), "[OPTIONS]")
	s.flagSet.StringVar(&s.opts.ProjectID, "p", "", "GCP Project ID")
	s.flagSet.StringVar(&s.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	s.flagSet.StringVar(&s.opts.SubscriptionID, "s", "", "Cloud Pub/Sub Subscription ID")
	return s
}

func (c *Sub) Name() string {
	return "sub"
}

func (c *Sub) Run(args []string) error {
	subcmd := args[1]
	c.flagSet.Parse(args[2:])

	if c.opts.ProjectID == "" {
		return errors.New("projectID must be provided")
	}

	s, err := tps.NewSubscriber(c.opts.ProjectID, c.opts.TopicID, c.opts.SubscriptionID)
	if err != nil {
		return err
	}

	ctx := context.Background()
	switch subcmd {
	case "create":
		if c.opts.TopicID == "" {
			return errors.New("topicID must be provided")
		}
		if c.opts.SubscriptionID == "" {
			return errors.New("subscriptionID must be provided")
		}

		s, err := s.CreateSubscription(ctx)
		fmt.Printf("subscription created: %v", s)
		if err != nil {
			return err
		}
	case "receive":
		if err := s.ReceiveSampleMessages(ctx); err != nil {
			return err
		}
	default:
		return errors.Errorf("subcommand is not supported: %s", subcmd)
	}
	return nil
}
