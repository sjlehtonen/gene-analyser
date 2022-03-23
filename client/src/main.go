package main

import (
	"github.com/sjlehtonen/gene-analyser/client/gocli"
	"github.com/sjlehtonen/gene-analyser/client/rabbitmq"
)

const PUBLISHER_ADDRESS = "amqp://guest:guest@localhost:5672/"

var CurrentPublisher *rabbitmq.Publisher

func AnalyseGeneHandler(cli *gocli.CLI, args []string) {
	if len(args) == 0 {
		cli.PrintError("argument[1] gene name is required for analyse-gene command")
		return
	}
	geneName := args[0]
	CurrentPublisher.SendMessage(geneName)
}

func main() {
	CurrentPublisher = rabbitmq.CreatePublisher(PUBLISHER_ADDRESS)
	gocli.RegisterCommandHandler("analyse-gene", AnalyseGeneHandler)
	gocli.Run()
}
