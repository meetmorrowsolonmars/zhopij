FROM golang:1.20

WORKDIR /project

COPY api api/internal
COPY cmd cmd/
COPY internal internal/
COPY go.mod go.sum ./

RUN go build -o ./bin/quiz ./cmd/quiz

ENTRYPOINT /project/bin/quiz
