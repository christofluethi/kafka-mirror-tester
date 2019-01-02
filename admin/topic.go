package admin

// Package admin is used for kafka's admin api

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"gitlab.appsflyer.com/rantav/kafka-mirror-tester/types"
)

// MustCreateTopic creates a new topic with the specified number of partitions.
// If the topic already exists, fails silently
// On error - simply panics
func MustCreateTopic(ctx context.Context, brokers types.Brokers, topic types.Topic, partitions, replicas int) {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": string(brokers)})
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}
	defer a.Close()

	res, err := a.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{
			{
				Topic:             string(topic),
				NumPartitions:     partitions,
				ReplicationFactor: replicas,
			},
		})
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	log.Infof("Topic create result: %v", res)
}