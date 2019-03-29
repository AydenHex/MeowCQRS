package main

import (
	"cqrs/db"
	"cqrs/event"
	"cqrs/search"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"time"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticsearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/meows", listMeowsHandler).
		Methods("GET")
	router.HandleFunc("/search", searchMeowsHandler).
		Methods("GET")
	return
}

func main() {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	// postgres connexions
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)
		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	// elk connexions
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := search.NewElastic(fmt.Sprintf("http://%s", conf.ElasticsearchAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		search.SetRepository(es)
		return nil
	})

	// nats connexion
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := event.NewNats(fmt.Sprintf("nats://%s", conf.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

	//router call
	router := NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
