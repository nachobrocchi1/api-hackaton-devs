package main

import (
	"api-hackaton-devs/client"
	"api-hackaton-devs/endpoint"
	"api-hackaton-devs/handler"
	"api-hackaton-devs/repository"
	"api-hackaton-devs/service"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/lib/pq"
	"gopkg.in/robfig/cron.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBERROR          = "Database Error"
	URI_USER_SERVICE = "https://randomuser.me/api?results=10"
	DBNAME           = "postgres"
	DBUSER           = "postgres"
	DBPASS           = "docker"
	DBPORT           = "5432"
	DBHOST           = "localhost"
)

var (
	logger     = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	devsClient = client.NewDevsClient(URI_USER_SERVICE, time.Duration(10)*time.Second, logger)
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// DataBase configuration
	connStr := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, DBHOST, DBPORT, DBUSER, DBPASS, DBNAME, "disable")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		level.Error(logger).Log(DBERROR, err)
	}
	db.AutoMigrate(&repository.Hackaton{}, &repository.Developer{})

	// Database population
	populateDB(db)

	// Initialization
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ep := endpoint.MakeServiceEndpoint(svc)
	ep = endpoint.LoggingEndpointMiddleware(logger)(ep)
	handler := handler.NewHTTPHandler(ep)

	// CRON every 5 minutes
	cron := cron.New()
	cron.AddFunc("@every 5m", func() {
		repo.InsertHackaton(getHackaton())
		level.Info(logger).Log("CRON", "success")
	})
	cron.Start()

	// start
	http.Handle("/hackaton", handler)
	level.Info(logger).Log("port", ":8080")
	level.Info(logger).Log("Error serving", http.ListenAndServe(":8080", nil))

}

func populateDB(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		h := getHackaton()
		db.Create(h)
	}
}
func getHackaton() *repository.Hackaton {
	h := new(repository.Hackaton)
	hID := rand.Intn(100) * rand.Intn(10)
	h.Name = fmt.Sprintf("Hackaton %d", hID)
	h.Devs = getHackatonDevs()
	return h
}

func getHackatonDevs() []repository.Developer {
	devs := make([]repository.Developer, 10)
	dev, err := devsClient.Call(context.Background(), nil)
	if err != nil {
		level.Error(logger).Log("population", "Cannot retreive devs")
	} else {
		for i := 0; i < 10; {
			devs[i] = repository.Developer{
				Position: i + 1,
				Name:     dev.Results[i].Name.First,
				LastName: dev.Results[i].Name.Last,
			}
			i++

		}
	}
	return devs
}
