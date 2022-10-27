package main

import (
	"log"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/config"
	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/service"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sendgrid/sendgrid-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		hclog.Default().Error(msg, "error", err)
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	l := hclog.Default()
	godotenv.Load(".env")
	l.Info("Loaded env vars")
	l.Info("Starting message microservice")
	conf := config.New()

	emailClient := sendgrid.NewSendClient(conf.SENDGRID_API_KEY)
	emailService := service.NewEmailService(emailClient)

	stop := make(chan bool)

	emailService.StartConsuming()

	<-stop

}
