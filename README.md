# pse
CLI for Cloud Pub/Sub emulator


# setup

## emulator

see https://cloud.google.com/pubsub/docs/emulator


### install

```bash
$ gcloud components install pubsub-emulator
$ gcloud components update
```

### start

```bash
$ gcloud beta emulators pubsub start
```

### envvar

```bash
$ $(gcloud beta emulators pubsub env-init)

# PUBSUB_EMULATOR_HOST is set. This switches Pub/Sub host address in cloud.google.com/go/pubsub.
# https://github.com/googleapis/google-cloud-go/blob/master/pubsub/pubsub.go#L60
```

Now, you can use a Cloud Pub/Sub emulator with pse. 

### stop

`Control+C` to stop an emulator.

```bash
$ unset PUBSUB_EMULATOR_HOST
```


# usage

1. go run cmd/pse.go create-topic -p testProject -t testTopic
2. go run cmd/pse.go create-sub -p testProject -t testTopic -s testSubscription
3. go run cmd/pse.go publish-sample -p testProject -t testTopic
4. go run cmd/pse.go receive-sample -p testProject -s testSubscription

You can show help and options.

```bash
$ go run cmd/pse.go -h

Usage: pse COMMAND [OPTIONS]
  -h --help
    Show commands
  -v --version
    Show version
Commands:
  create-topic  Create topic
  delete-topic  Delete topic
  publish-sample  Publish sample messages
  create-sub  Create subscription
  receive-sample  Receive sample messages as a subscriber process

See more details:
  pse COMMAND -h
```
