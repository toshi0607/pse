package subcmd

import "errors"

type (
	Command interface {
		Name() string
		Run([]string) error
	}

	CommandRepository interface {
		Find(string) (Command, error)
		Commands() []Command
	}

	repository struct {
		commands   []Command
		commandMap map[string]Command
	}
)

func Repository() CommandRepository {
	repo := &repository{
		commandMap: make(map[string]Command),
	}
	repo.commands = []Command{
		NewPub(),
		NewSub(),
	}
	for _, c := range repo.commands {
		repo.commandMap[c.Name()] = c
	}
	return repo
}

func (repo *repository) Find(name string) (Command, error) {
	c, ok := repo.commandMap[name]
	if !ok {
		return nil, errors.New("command is missing")
	}
	return c, nil
}

func (repo *repository) Commands() []Command {
	return repo.commands[:]
}
