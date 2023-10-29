package producers

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/message"
)

type AnswerEventsProducer struct {
	topic    string
	producer sarama.SyncProducer
	options  protojson.MarshalOptions
}

func NewAnswerEventsProducer(topic string, producer sarama.SyncProducer, options protojson.MarshalOptions) *AnswerEventsProducer {
	return &AnswerEventsProducer{
		topic:    topic,
		producer: producer,
		options:  options,
	}
}

func (p *AnswerEventsProducer) SendCreateEvent(answer *domain.Answer) error {
	msg, err := p.options.Marshal(&message.V1CreateAnswerEvent{
		Meta: &message.Meta{
			MessageId: uuid.New().String(),
			Version:   message.MessageVersion_V1_0_0,
			Event:     message.EventType_CREATE,
		},
		Event: &message.V1CreateAnswerEvent_Event{
			QuizId:   answer.QuizID,
			AuthorId: answer.AuthorID,
		},
	})
	if err != nil {
		return err
	}

	_, _, err = p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(fmt.Sprintf("author_id:%d,quiz_id:%d", answer.AuthorID, answer.QuizID)),
		Value: sarama.ByteEncoder(msg),
	})
	return err
}
