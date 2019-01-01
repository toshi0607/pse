package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/mock/pubsub"
)

func TestCreateTopic_NewCreateTopic(t *testing.T) {
	t.Run("create CreateTopic successfully", func(t *testing.T) {
		s := NewCreateTopic()
		if s == nil {
			t.Errorf("reateTopic should be created")
		}
	})
}

func TestCreateTopic_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "create-topic"

		s := NewCreateTopic()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestCreateTopic_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Create topic"

		s := NewCreateTopic()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestCreateTopic_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fpub := mock_pubsub.NewMockPublisher(ctrl)

	tests := map[string]struct {
		init      func()
		args      []string
		wantError bool
	}{
		"normal": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().CreateTopic(gomock.Any(), "testt").Return(nil, nil)
			}, []string{"create-topic", "-p", "testp", "-t", "testt"}, false,
		},
		"with help option": {
			func() {},
			[]string{"create-topic", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"create-topic", "-t", "testt"}, true,
		},
		"without topic ID": {
			func() {},
			[]string{"create-topic", "-p", "testp"}, true,
		},
		"with Publisher.Init error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"create-topic", "-p", "testp", "-t", "testt"}, true,
		},
		"with Publisher.CreateTopic error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().CreateTopic(gomock.Any(), "testt").Return(nil, errors.New("test error"))
			},
			[]string{"create-topic", "-p", "testp", "-t", "testt"}, true,
		},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			te.init()
			cmd := NewCreateTopic()
			cmd.pub = fpub

			if err := cmd.Run(te.args); err != nil {
				if !te.wantError {
					t.Errorf("want no error, got: %v", err)
				}
			}
		})
	}
}
