package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sjlehtonen/gene-analyser/server/integration"
	"github.com/sjlehtonen/gene-analyser/server/rabbitmq"
	"github.com/sjlehtonen/gene-analyser/server/report"
)

var CurrentConsumer *rabbitmq.Consumer

const REPORTS_PATH = "/reports/"
const WEBSERVER_PORT = ":8080"
const RABBIT_MQ_ADDRESS = "amqp://guest:guest@localhost:5672/"

func InitHttpHandlers() {
	http.Handle(REPORTS_PATH, http.StripPrefix(REPORTS_PATH, http.FileServer(http.Dir("../reports"))))
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func HandleReceiveMessage(body []byte) {
	geneName := string(body)
	fmt.Printf("server: procesing gene %s...\n", geneName)
	dna, description, err := integration.GetDNAAndDescriptionForGene(geneName)
	if err != nil {
		handleError(err)
		return
	}
	geneReport := report.GenerateReportForGene(geneName, dna, description)
	geneReport.WriteReport()
	fmt.Printf("server: finished processing gene %s\n", geneName)
}

func main() {
	fmt.Println("-- Starting server --")
	go rabbitmq.CreateConsumer(RABBIT_MQ_ADDRESS).ReceiveMessages(HandleReceiveMessage)
	InitHttpHandlers()
	log.Fatal(http.ListenAndServe(WEBSERVER_PORT, nil))
}
