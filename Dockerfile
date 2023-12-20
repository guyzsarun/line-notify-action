FROM golang:1.21.4-bullseye

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /line-notify-action
CMD ["/line-notify-action"]