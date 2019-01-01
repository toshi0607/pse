package subcmd

import "testing"

func TestRepository(t *testing.T) {
	r := Repository()
	if len(r.Commands()) != 6 {
		t.Errorf("got: %d, want: 5", len(r.Commands()))
	}
}

func TestRepository_Commands(t *testing.T) {
	repo := &repository{
		commandMap: make(map[string]Command),
	}
	repo.commands = []Command{
		NewCreateTopic(),
		NewDeleteTopic(),
	}

	cs := repo.Commands()
	if len(cs) != 2 {
		t.Errorf("got: %d, want: 2", len(cs))
	}
}

func TestRepository_Find(t *testing.T) {
	repo := &repository{
		commandMap: make(map[string]Command),
	}
	repo.commands = []Command{
		NewCreateTopic(),
		NewDeleteTopic(),
	}
	for _, c := range repo.commands {
		repo.commandMap[c.Name()] = c
	}

	tests := map[string]struct {
		findName  string
		want      string
		wantError bool
	}{
		"normal1":         {"create-topic", "create-topic", false},
		"normal2":         {"delete-topic", "delete-topic", false},
		"command missing": {"hoge", "", true},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := repo.Find(te.findName)
			if te.wantError {
				if err == nil {
					t.Error("want error, got nothing")
				}
			} else {
				if err != nil {
					t.Errorf("want no error, got: %v", err)
				}

				if c.Name() != te.want {
					t.Errorf("want %s, got: %s", te.want, c.Name())
				}
			}
		})
	}
}
