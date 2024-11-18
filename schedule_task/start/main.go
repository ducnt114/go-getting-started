package main

import (
	"context"
	"go-getting-started/schedule_task"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func main() {
	// Connect to Temporal server
	c, err := client.Dial(client.Options{
		HostPort: "103.20.96.166:17233",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	ctx := context.Background()

	// Define the schedule options
	scheduleSpec := client.ScheduleSpec{
		Intervals: []client.ScheduleIntervalSpec{
			{Every: time.Minute}, // Execute every minute
		},
	}

	// Define the workflow to execute on the schedule
	workflowID := "ScheduledWorkflow"

	// Combine schedule and workflow options
	scheduleID := "example-schedule-2"
	scheduleOptions := client.ScheduleOptions{
		ID:   scheduleID,
		Spec: scheduleSpec,
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule_task.ScheduledWorkflow,
			TaskQueue: schedule_task.ScheduleTaskQueue,
		},
	}

	// Create the schedule
	_, err = c.ScheduleClient().Create(ctx, scheduleOptions)
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}

	log.Println("Schedule created successfully with ID:", scheduleID)
}
