package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/pubsub"
)

func TestDeleteTopic_NewDeleteTopic(t *testing.T) {
	t.Run("create DeleteTopic successfully", func(t *testing.T) {
		s := NewDeleteTopic()
		if s == nil {
			t.Errorf("DeleteTopic should be created")
		}
	})
}

func TestDeleteTopic_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "delete-topic"

		s := NewDeleteTopic()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestDeleteTopic_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Delete topic"

		s := NewDeleteTopic()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestDeleteTopic_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fpub := pubsub.NewMockPublisher(ctrl)

	tests := map[string]struct {
		init      func()
		args      []string
		wantError bool
	}{
		"normal": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().DeleteTopic(gomock.Any(), "testt").Return(nil)
			}, []string{"delete-topic", "-p", "testp", "-t", "testt"}, false,
		},
		"with help option": {
			func() {},
			[]string{"delete-topic", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"delete-topic", "-t", "testt"}, true,
		},
		"without topic ID": {
			func() {},
			[]string{"delete-topic", "-p", "testp"}, true,
		},
		"with Publisher.Init error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"delete-topic", "-p", "testp", "-t", "testt"}, true,
		},
		"with Publisher.DeleteTopic error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().DeleteTopic(gomock.Any(), "testt").Return(errors.New("test error"))
			},
			[]string{"delete-topic", "-p", "testp", "-t", "testt"}, true,
		},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			te.init()
			cmd := NewDeleteTopic()
			cmd.pub = fpub

			if err := cmd.Run(te.args); err != nil {
				if !te.wantError {
					t.Errorf("want no error, got: %v", err)
				}
			}
		})
	}
}
