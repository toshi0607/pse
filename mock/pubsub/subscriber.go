// Code generated by MockGen. DO NOT EDIT.
// Source: pubsub/subscriber.go

// Package mock_pubsub is a generated GoMock package.
package mock_pubsub

import (
	pubsub "cloud.google.com/go/pubsub"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSubscriber is a mock of Subscriber interface
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// CreateSubscription mocks base method
func (m *MockSubscriber) CreateSubscription(ctx context.Context, subscriptionID, topicID string) (*pubsub.Subscription, error) {
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, subscriptionID, topicID)
	ret0, _ := ret[0].(*pubsub.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription
func (mr *MockSubscriberMockRecorder) CreateSubscription(ctx, subscriptionID, topicID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockSubscriber)(nil).CreateSubscription), ctx, subscriptionID, topicID)
}

// ReceiveSampleMessages mocks base method
func (m *MockSubscriber) ReceiveSampleMessages(ctx context.Context, subscriptionID string) error {
	ret := m.ctrl.Call(m, "ReceiveSampleMessages", ctx, subscriptionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveSampleMessages indicates an expected call of ReceiveSampleMessages
func (mr *MockSubscriberMockRecorder) ReceiveSampleMessages(ctx, subscriptionID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveSampleMessages", reflect.TypeOf((*MockSubscriber)(nil).ReceiveSampleMessages), ctx, subscriptionID)
}

// Init mocks base method
func (m *MockSubscriber) Init(ctx context.Context, projectID string) error {
	ret := m.ctrl.Call(m, "Init", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockSubscriberMockRecorder) Init(ctx, projectID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSubscriber)(nil).Init), ctx, projectID)
}
