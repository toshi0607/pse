# pse
CLI for Cloud Pub/Sub emulator

# usage

1. go run cmd/pse.go create-topic -p testProject -t testTopic
2. go run cmd/pse.go create-sub -p testProject -t testTopic -s testSubscription
3. go run cmd/pse.go publish-sample -p testProject -t testTopic
4. go run cmd/pse.go receive-sample -p testProject -s testSubscription
