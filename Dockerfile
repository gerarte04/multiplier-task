FROM golang:1.24.2-alpine3.21

ENV RTP=0.95

COPY . .
RUN go build -o main main.go

RUN chmod a+x inspector_linux_amd64

CMD go run . -rtp=${RTP}