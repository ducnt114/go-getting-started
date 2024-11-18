package schedule_task

import "go.temporal.io/sdk/workflow"

// Workflow definition
func ScheduledWorkflow(ctx workflow.Context) error {
	// Log the workflow execution
	workflow.GetLogger(ctx).Info("Scheduled workflow executed at", "time", workflow.Now(ctx))
	return nil
}
