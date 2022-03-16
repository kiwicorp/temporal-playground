package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kiwicorp/temporal-playground/internal/helloworld"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: helloworld.TransferMoneyTaskQueue,
	}
	transferDetails := helloworld.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, helloworld.TransferMoney, transferDetails)
	if err != nil {
		log.Fatalln("error starting TransferMoney workflow", err)
	}
	printResults(transferDetails, we.GetID(), we.GetRunID())
}

func printResults(transferDetails helloworld.TransferDetails, workflowID, runID string) {
	log.Printf(
		"\nTransfer of $%f from account %s to account %s is processing. ReferenceID: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
