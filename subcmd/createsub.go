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
	CreateSub struct {
		flagSet *flag.FlagSet
		opts    CreateSubConfig
	}

	CreateSubConfig struct {
		ProjectID, TopicID, SubscriptionID string
		ShowHelp                           bool
	}
)

func NewCreateSub() *CreateSub {
	c := &CreateSub{}
	c.flagSet = tf.New(c.Name(), "[OPTIONS]")
	c.flagSet.StringVar(&c.opts.ProjectID, "p", "", "GCP Project ID")
	c.flagSet.StringVar(&c.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	c.flagSet.StringVar(&c.opts.SubscriptionID, "s", "", "Cloud Pub/Sub Subscription ID")
	c.flagSet.BoolVar(&c.opts.ShowHelp, "h", false, "Show command usage")
	return c
}

func (c *CreateSub) Name() string {
	return "create-sub"
}

func (c *CreateSub) Summary() string {
	return "Create subscription"
}

func (c *CreateSub) Usage() {
	c.flagSet.Usage()
}

func (c *CreateSub) Run(args []string) error {
	c.flagSet.Parse(args[1:])
	if c.opts.ShowHelp {
		c.Usage()
		return nil
	}

	if c.opts.ProjectID == "" {
		return errors.New("projectID must be provided")
	}
	if c.opts.TopicID == "" {
		return errors.New("topicID must be provided")
	}
	if c.opts.SubscriptionID == "" {
		return errors.New("subscriptionID must be provided")
	}

	s, err := tps.NewSubscriber(c.opts.ProjectID, c.opts.TopicID, c.opts.SubscriptionID)
	if err != nil {
		return err
	}

	ctx := context.Background()
	sub, err := s.CreateSubscription(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("subscription created: %v", sub)

	return nil
}
