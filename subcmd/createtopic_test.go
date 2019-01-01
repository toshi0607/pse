package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/mock/pubsub"
)

func TestCreateTopic_NewCreateSub(t *testing.T) {
	t.Run("create CreateSub successfully", func(t *testing.T) {
		s := NewCreateTopic()
		if s == nil {
			t.Errorf("CreateSub should be created")
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
			}, []string{"create-sub", "-p", "testp", "-t", "testt"}, false,
		},
		"with help option": {
			func() {},
			[]string{"create-sub", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"create-sub", "-t", "testt"}, true,
		},
		"without topic ID": {
			func() {},
			[]string{"create-sub", "-p", "testp"}, true,
		},
		"with Subscriber.Init error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"create-sub", "-p", "testp", "-t", "testt"}, true,
		},
		"with Subscriber.CreateSubscription error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().CreateTopic(gomock.Any(), "testt").Return(nil, errors.New("test error"))
			},
			[]string{"create-sub", "-p", "testp", "-t", "testt"}, true,
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
