package main

import (
	"go-getting-started/money_transfer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, money_transfer.MoneyTransferTaskQueueName, worker.Options{})

	// This worker_workflow hosts both Workflow and Activity functions.
	w.RegisterWorkflow(money_transfer.MoneyTransfer)
	w.RegisterActivity(money_transfer.Withdraw)
	w.RegisterActivity(money_transfer.Deposit)
	w.RegisterActivity(money_transfer.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
