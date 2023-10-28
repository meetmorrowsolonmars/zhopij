package consumers

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type EchoConsumer struct {
	logger *zap.SugaredLogger
}

func NewEchoConsumer(logger *zap.SugaredLogger) *EchoConsumer {
	return &EchoConsumer{
		logger: logger,
	}
}

func (e *EchoConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (e *EchoConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (e *EchoConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				e.logger.Infof("msg channel was closed")
				return nil
			}

			e.logger.Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic)
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
