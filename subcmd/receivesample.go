package subcmd

import (
	"context"
	"flag"

	"github.com/pkg/errors"
	tf "github.com/toshi0607/pse/flag"
	tps "github.com/toshi0607/pse/pubsub"
)

type (
	ReceiveSample struct {
		flagSet *flag.FlagSet
		opts    ReceiveSampleConfig
		sub     tps.Subscriber
	}

	ReceiveSampleConfig struct {
		ProjectID, TopicID, SubscriptionID string
		ShowHelp                           bool
	}
)

func NewReceiveSample() *ReceiveSample {
	c := &ReceiveSample{sub: tps.NewSubscriber()}
	c.flagSet = tf.New(c.Name(), "[OPTIONS]")
	c.flagSet.StringVar(&c.opts.ProjectID, "p", "", "GCP Project ID")
	c.flagSet.StringVar(&c.opts.SubscriptionID, "s", "", "Cloud Pub/Sub Subscription ID")
	c.flagSet.BoolVar(&c.opts.ShowHelp, "h", false, "Show command usage")
	return c
}

func (c *ReceiveSample) Name() string {
	return "receive-sample"
}

func (c *ReceiveSample) Summary() string {
	return "Receive sample messages as a subscriber process"
}

func (c *ReceiveSample) Usage() {
	c.flagSet.Usage()
}

func (c *ReceiveSample) Run(args []string) error {
	c.flagSet.Parse(args[1:])
	if c.opts.ShowHelp {
		c.Usage()
		return nil
	}

	if c.opts.ProjectID == "" {
		return errors.New("projectID must be provided")
	}
	if c.opts.SubscriptionID == "" {
		return errors.New("subscriptionID must be provided")
	}

	ctx := context.Background()
	if err := c.sub.Init(ctx, c.opts.ProjectID); err != nil {
		return err
	}
	if err := c.sub.ReceiveSampleMessages(ctx, c.opts.SubscriptionID); err != nil {
		return err
	}

	return nil
}
