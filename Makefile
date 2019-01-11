.PHONY: gen_mock
gen_mock:
	@go get github.com/golang/mock/mockgen
	mockgen -source pubsub/publisher.go -destination pubsub/mock_publisher.go -package pubsub
	mockgen -source pubsub/subscriber.go -destination pubsub/mock_subscriber.go -package pubsub
