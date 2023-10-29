package consumers

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/domain"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pb/api/v1/message"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/services"
)

type AnswerEventsConsumer struct {
	service *services.AnswerService
	logger  *zap.SugaredLogger
	options protojson.UnmarshalOptions
}

func NewAnswerEventsConsumer(service *services.AnswerService, logger *zap.SugaredLogger, options protojson.UnmarshalOptions) *AnswerEventsConsumer {
	return &AnswerEventsConsumer{
		service: service,
		logger:  logger,
		options: options,
	}
}

func (c *AnswerEventsConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AnswerEventsConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AnswerEventsConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				c.logger.Info("msg channel was closed")
				return nil
			}

			// TODO: check topic name

			var baseEvent message.BaseEvent
			err := c.options.Unmarshal(msg.Value, &baseEvent)
			if err != nil {
				// restart consumer
				c.logger.Errorf("can't unmarshal answer event: %s", err)
				return err
			}

			// TODO: add fabric by event type
			// create answer
			if baseEvent.Meta.Version == message.MessageVersion_V1_0_0 && baseEvent.Meta.Event == message.EventType_CREATE {
				var createEvent message.V1CreateAnswerEvent
				err = c.options.Unmarshal(msg.Value, &createEvent)
				if err != nil {
					// restart consumer
					c.logger.Errorf("can't unmarshal answer event: %s", err)
					return err
				}

				_, err = c.service.Create(session.Context(), &domain.Answer{
					QuizID:   createEvent.Event.QuizId,
					AuthorID: createEvent.Event.AuthorId,
				})
				if err != nil {
					// restart consumer
					c.logger.Errorf("can't create answer: %s", err)
					return err
				}
			}

			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
