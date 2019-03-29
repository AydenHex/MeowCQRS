package main

import (
	"cqrs/event"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"time"
)

type Config struct {
	NatsAdresss string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}
}
