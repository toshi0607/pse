package subcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/toshi0607/pse/mock/pubsub"
)

func TestPublishSample_NewPublishSample(t *testing.T) {
	t.Run("create PublishSample successfully", func(t *testing.T) {
		s := NewPublishSample()
		if s == nil {
			t.Errorf("PublishSample should be created")
		}
	})
}

func TestPublishSample_Name(t *testing.T) {
	t.Run("show a sub command's name successfully", func(t *testing.T) {
		want := "publish-sample"

		s := NewPublishSample()
		got := s.Name()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestPublishSample_Summary(t *testing.T) {
	t.Run("show a sub command's summary successfully", func(t *testing.T) {
		want := "Publish sample messages"

		s := NewPublishSample()
		got := s.Summary()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestPublishSample_Run(t *testing.T) {
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
				fpub.EXPECT().PublishSampleMessage(gomock.Any(), "testt").Return(nil)
			}, []string{"publish-sample", "-p", "testp", "-t", "testt"}, false,
		},
		"with help option": {
			func() {},
			[]string{"publish-sample", "-h"}, false,
		},
		"without project ID": {
			func() {},
			[]string{"publish-sample", "-t", "testt"}, true,
		},
		"without topic ID": {
			func() {},
			[]string{"publish-sample", "-p", "testp"}, true,
		},
		"with Publisher.Init error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(errors.New("test error"))
			},
			[]string{"publish-sample", "-p", "testp", "-t", "testt"}, true,
		},
		"with Publisher.PublishSampleMessage error": {
			func() {
				fpub.EXPECT().Init(gomock.Any(), "testp").Return(nil)
				fpub.EXPECT().PublishSampleMessage(gomock.Any(), "testt").Return(errors.New("test error"))
			},
			[]string{"publish-sample", "-p", "testp", "-t", "testt"}, true,
		},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			te.init()
			cmd := NewPublishSample()
			cmd.pub = fpub

			if err := cmd.Run(te.args); err != nil {
				if !te.wantError {
					t.Errorf("want no error, got: %v", err)
				}
			}
		})
	}
}
