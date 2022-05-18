# Nats-PubSub
A simple PubSub demo using NATS

Run [NATS](https://nats.io/) server with

`docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats-server -ti nats:latest`

run Subscriber
`go run sub/main.go`

run Publisher
`go run sub/main.go`