package main

import (
	"log"

	"github.com/kiwicorp/temporal-playground/internal/helloworld"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, helloworld.TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(helloworld.TransferMoney)
	w.RegisterActivity(helloworld.Withdraw)
	w.RegisterActivity(helloworld.Deposit)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
