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
	DeleteTopic struct {
		flagSet *flag.FlagSet
		opts    CreateTopicConfig
		pub     tps.Publisher
	}

	DeleteTopicConfig struct {
		ProjectID, TopicID string
		ShowHelp           bool
	}
)

func NewDeateTopic() *DeleteTopic {
	c := &DeleteTopic{pub: tps.NewPublisher()}
	c.flagSet = tf.New(c.Name(), "[OPTIONS]")
	c.flagSet.StringVar(&c.opts.ProjectID, "p", "", "GCP Project ID")
	c.flagSet.StringVar(&c.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	c.flagSet.BoolVar(&c.opts.ShowHelp, "h", false, "Show command usage")
	return c
}

func (c *DeleteTopic) Name() string {
	return "delete-topic"
}

func (c *DeleteTopic) Summary() string {
	return "Delete topic"
}

func (c *DeleteTopic) Usage() {
	c.flagSet.Usage()
}

func (c *DeleteTopic) Run(args []string) error {
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

	if err := c.pub.DeleteTopic(ctx, c.opts.TopicID); err != nil {
		return err
	}
	fmt.Printf("topic deleted: %v", c.opts.TopicID)

	return nil
}
