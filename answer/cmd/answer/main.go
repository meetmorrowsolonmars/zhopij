package main

import (
	"log"

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
}
