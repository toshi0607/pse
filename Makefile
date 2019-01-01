.PHONY: gen_mock
gen_mock:
	@go get github.com/golang/mock/mockgen
	mockgen -source pubsub/publisher.go -destination mock/pubsub/publisher.go
	mockgen -source pubsub/subscriber.go -destination mock/pubsub/subscriber.go
