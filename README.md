# pse
CLI for Cloud Pub/Sub emulator

# usage

1. go run main.go pub create -p testProject -t testTopic
2. go run main.go sub create -p testProject -t testTopic -s testSubscription
3. go run main.go pub publish -t testProject -t testTopic
4. go run main.go sub receive -t testProject -s testSubscription
