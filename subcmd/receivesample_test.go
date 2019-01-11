package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/pubsub"
)

func TestReceiveSample_NewReceiveSample(t *testing.T) {
	t.Run("create ReceiveSample successfully", func(t *testing.T) {
		s := NewReceiveSample()
		if s == nil {
			t.Errorf("ReceiveSample should be created")
		}
	})
}

func TestReceiveSample_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "receive-sample"

		s := NewReceiveSample()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestReceiveSample_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Receive sample messages as a subscriber process"

		s := NewReceiveSample()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestReceiveSample_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fsub := pubsub.NewMockSubscriber(ctrl)

	tests := map[string]struct {
		init      func()
		args      []string
		wantError bool
	}{
		"normal": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fsub.EXPECT().ReceiveSampleMessages(gomock.Any(), "tests").Return(nil)
			}, []string{"receive-sample", "-p", "testp", "-s", "tests"}, false,
		},
		"with help option": {
			func() {},
			[]string{"receive-sample", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"receive-sample", "-s", "tests"}, true,
		},
		"without subscription ID": {
			func() {},
			[]string{"receive-sample", "-p", "testt"}, true,
		},
		"with Subscriber.Init error": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"receive-sample", "-p", "testp", "-s", "tests"}, true,
		},
		"with Subscriber.ReceiveSamplescription error": {
			func() {
				fsub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fsub.EXPECT().ReceiveSampleMessages(gomock.Any(), "tests").Return(errors.New("test error"))
			},
			[]string{"receive-sample", "-p", "testp", "-s", "tests"}, true,
		},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			te.init()
			cmd := NewReceiveSample()
			cmd.sub = fsub

			if err := cmd.Run(te.args); err != nil {
				if !te.wantError {
					t.Errorf("want no error, got: %v", err)
				}
			}
		})
	}
}
