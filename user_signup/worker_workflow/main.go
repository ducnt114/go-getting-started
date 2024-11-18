package main

import (
	"go-getting-started/user_signup"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "103.20.96.166:17233",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker_workflow hosts both Workflow and Activity functions
	w := worker.New(c, user_signup.UserSignupTaskQueue, worker.Options{})

	w.RegisterWorkflow(user_signup.UserSignupWorkflow)

	//w.RegisterActivity(app.ComposeGreeting)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
