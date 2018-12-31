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
	CreateTopic struct {
		flagSet *flag.FlagSet
		opts    CreateTopicConfig
		pub     tps.Publisher
	}

	CreateTopicConfig struct {
		ProjectID, TopicID string
		ShowHelp           bool
	}
)

func NewCreateTopic() *CreateTopic {
	c := &CreateTopic{pub: tps.NewPublisher()}
	c.flagSet = tf.New(c.Name(), "[OPTIONS]")
	c.flagSet.StringVar(&c.opts.ProjectID, "p", "", "GCP Project ID")
	c.flagSet.StringVar(&c.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	c.flagSet.BoolVar(&c.opts.ShowHelp, "h", false, "Show command usage")
	return c
}

func (c *CreateTopic) Name() string {
	return "create-topic"
}

func (c *CreateTopic) Summary() string {
	return "Create topic"
}

func (c *CreateTopic) Usage() {
	c.flagSet.Usage()
}

func (c *CreateTopic) Run(args []string) error {
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

	ctx := context.Background()
	if err := c.pub.Init(ctx, c.opts.ProjectID); err != nil {
		return err
	}

	t, err := c.pub.CreateTopic(ctx, c.opts.TopicID)
	if err != nil {
		return err
	}
	fmt.Printf("topic created: %v", t)

	return nil
}
