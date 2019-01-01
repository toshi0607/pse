package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/mock/pubsub"
)

func TestCreateSub_NewCreateSub(t *testing.T) {
	t.Run("create CreateSub successfully", func(t *testing.T) {
		s := NewCreateSub()
		if s == nil {
			t.Errorf("CreateSub should be created")
		}
	})
}

func TestCreateSub_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "create-sub"

		s := NewCreateSub()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestCreateSub_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Create subscription"

		s := NewCreateSub()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestCreateSub_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fsub := mock_pubsub.NewMockSubscriber(ctrl)

	tests := map[string]struct {
		init      func()
		args      []string
		wantError bool
	}{
		"normal": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fsub.EXPECT().CreateSubscription(gomock.Any(), "tests", "testt").Return(nil, nil)
			}, []string{"create-sub", "-p", "testp", "-t", "testt", "-s", "tests"}, false,
		},
		"with help option": {
			func() {},
			[]string{"create-sub", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"create-sub", "-t", "testt", "-s", "tests"}, true,
		},
		"without topic ID": {
			func() {},
			[]string{"create-sub", "-p", "testp", "-s", "tests"}, true,
		},
		"without subscription ID": {
			func() {},
			[]string{"create-sub", "-p", "testp", "-t", "testt"}, true,
		},
		"with Subscriber.Init error": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"create-sub", "-p", "testp", "-t", "testt", "-s", "tests"}, true,
		},
		"with Subscriber.CreateSubscription error": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fsub.EXPECT().CreateSubscription(gomock.Any(), "tests", "testt").Return(nil, errors.New("test error"))
			},
			[]string{"create-sub", "-p", "testp", "-t", "testt", "-s", "tests"}, true,
		},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			te.init()
			cmd := NewCreateSub()
			cmd.sub = fsub

			if err := cmd.Run(te.args); err != nil {
				if !te.wantError {
					t.Errorf("want no error, got: %v", err)
				}
			}
		})
	}
}
