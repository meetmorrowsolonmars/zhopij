package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/message"
)

func main() {
	log.Println("Hello, World from Answer!")

	msg := &message.V1CreateAnswerEvent{
		Meta: &message.Meta{
			MessageId: uuid.New().String(),
			Version:   message.MessageVersion_V1_0_0,
			Event:     message.EventType_CREATE,
		},
		Event: &message.V1CreateAnswerEvent_Event{
			QuizId: 10,
			UserId: 200,
		},
	}

	payload, err := protojson.MarshalOptions{
		Multiline:      true,
		UseProtoNames:  true,
		UseEnumNumbers: false,
	}.Marshal(msg)
	if err != nil {
		log.Fatalln("can't marshal message")
	}

	log.Println(string(payload))

	brokers := []string{"kafka:29092"}
	groupID := "answer"
	consumerGroupConfig := sarama.NewConfig()
	consumerGroupConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
		sarama.NewBalanceStrategyRange(),
	}
	// TODO: test it
	consumerGroupConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, consumerGroupConfig)
	if err != nil {
		log.Fatalf("can't create consumer group: %s", err)
	}

	ctx := context.Background()
	err = consumerGroup.Consume(ctx, []string{"answers_test"}, &handler{})
	if err != nil {
		log.Fatalf("can't start consumer group: %s", err)
	}
}

type handler struct {
}

func (h handler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h handler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Printf("msg channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic)
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
