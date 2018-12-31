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
	Pub struct {
		flagSet *flag.FlagSet
		opts    PubConfig
	}

	PubConfig struct {
		ProjectID, TopicID string
		ShowHelp           bool
	}
)

func NewPub() *Pub {
	p := &Pub{}
	p.flagSet = tf.New(p.Name(), "[OPTIONS]")
	p.flagSet.StringVar(&p.opts.ProjectID, "p", "", "GCP Project ID")
	p.flagSet.StringVar(&p.opts.TopicID, "t", "", "Cloud Pub/Sub Topic ID")
	p.flagSet.BoolVar(&p.opts.ShowHelp, "h", false, "Show command usage")
	return p
}

func (c *Pub) Name() string {
	return "pub"
}

func (c *Pub) Summary() string {
	return `
pse pub SUBCOMMAND [OPTIONS]
  Sub Commands:
    create-topic
      Create topic. Use -p & -t
    publish
      Publish sample message. Use -p & -t
`
}

func (c *Pub) Usage() {
	c.flagSet.Usage()
}

func (c *Pub) Run(args []string) error {
	subcmd := args[1]
	c.flagSet.Parse(args[2:])
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
	switch subcmd {
	case "create-topic":
		t, err := p.CreateTopic(ctx)
		fmt.Printf("topic created: %v", t)
		if err != nil {
			return err
		}
	case "delete-topic":
		err := p.DeleteTopic(ctx)
		fmt.Printf("topic deleted: %v", c.opts.TopicID)
		if err != nil {
			return err
		}
	case "publish":
		if err := p.PublishSampleMessage(ctx); err != nil {
			return err
		}
	default:
		return errors.Errorf("subcommand is not supported: %s", subcmd)
	}

	return nil
}
