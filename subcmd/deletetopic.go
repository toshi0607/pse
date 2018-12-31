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
	DeateTopic struct {
		flagSet *flag.FlagSet
		opts    CreateTopicConfig
	}

	DeateTopicConfig struct {
		ProjectID, TopicID string
		ShowHelp           bool
	}
)

func NewDeateTopic() *DeateTopic {
	c := &DeateTopic{}
	c.flagSet = tf.New(c.Name(), "[OPTIONS]")
	c.flagSet.StringVar(&c.opts.ProjectID, "p", "", "GCP Project ID")
	c.flagSet.StringVar(&c.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	c.flagSet.BoolVar(&c.opts.ShowHelp, "h", false, "Show command usage")
	return c
}

func (c *DeateTopic) Name() string {
	return "delete-topic"
}

func (c *DeateTopic) Summary() string {
	return "Delete topic"
}

func (c *DeateTopic) Usage() {
	c.flagSet.Usage()
}

func (c *DeateTopic) Run(args []string) error {
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

	p, err := tps.NewPublisher(c.opts.ProjectID, c.opts.TopicID)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := p.DeleteTopic(ctx); err != nil {
		return err
	}
	fmt.Printf("topic deleted: %v", c.opts.TopicID)

	return nil
}
