package producers

import (
	"github.com/IBM/sarama"
)

type AnswerEventsProducer struct {
	topic    string
	producer sarama.SyncProducer
}

func NewAnswerEventsProducer(topic string, producer sarama.SyncProducer) *AnswerEventsProducer {
	return &AnswerEventsProducer{
		topic:    topic,
		producer: producer,
	}
}

func (p *AnswerEventsProducer) SendCreateEvent(msg []byte) error {
	// TODO: use domain model
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(msg),
	})
	return err
}
