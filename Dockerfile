FROM golang:1.21.4-bullseye
LABEL   org.opencontainers.image.source=https://github.com/guyzsarun/line-notify-action \
        org.opencontainers.image.licenses=Apache-2.0 \
        org.opencontainers.image.authors=sarun.nunt@gmail.com


WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /line-notify-action

CMD ["/line-notify-action"]