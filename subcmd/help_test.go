package subcmd

import "testing"

func TestHelp_NewHelp(t *testing.T) {
	t.Run("create Help successfully", func(t *testing.T) {
		s := NewHelp()
		if s == nil {
			t.Errorf("Help should be created")
		}
	})
}

func TestHelp_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "help"

		s := NewHelp()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestHelp_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Show commands"

		s := NewHelp()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestHelp_Run(t *testing.T) {
	cmd := NewHelp()
	err := cmd.Run(nil)
	if err != nil {
		t.Errorf("want no error, got: %v", err)
	}
}
