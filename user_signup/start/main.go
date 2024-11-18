package main

import (
	"context"
	"fmt"
	"go-getting-started/user_signup"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
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

	resp, err := c.ListWorkflow(context.Background(),
		&workflowservice.ListWorkflowExecutionsRequest{})
	if err != nil {
		log.Fatalln("unable to list workflows", err)
		return
	}
	for _, ex := range resp.GetExecutions() {
		fmt.Println("name: " + ex.Type.GetName())
	}

	options := client.StartWorkflowOptions{
		ID:        "user-signup-workflow",
		TaskQueue: user_signup.UserSignupTaskQueue,
	}

	// Start the Workflow
	req := user_signup.UserSignupRequest{
		Username: "duc-test-1",
		Password: "random-pass",
		Email:    "error@gmail.com",
		//Email: "valid@gmail.com",
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, user_signup.UserSignupWorkflow, req)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
	var signupResp string
	err = we.Get(context.Background(), &signupResp)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(signupResp, we.GetID(), we.GetRunID())
}

func printResults(signupResp string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", signupResp)
}
