package user_signup

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func UserSignupWorkflow(ctx workflow.Context, req UserSignupRequest) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var saveDbResult string
	err := workflow.ExecuteActivity(ctx, SaveUserToDatabase, req).Get(ctx, &saveDbResult)
	if err != nil {
		return "", err
	}
	log.Printf("saveDbResult: %v\n\n", saveDbResult)

	_, err = workflow.AwaitWithTimeout(ctx, 30*time.Second, func() bool {
		// logic to verify user email
		if time.Now().Unix() > 1731920844 {
			return true
		}
		return false
	})
	if err != nil {
		return "", err
	}

	var sendEmailResult string
	err = workflow.ExecuteActivity(ctx, SendWelcomeEmail, req).Get(ctx, &sendEmailResult)
	if err != nil {
		return "", err
	}
	log.Printf("sendEmailResult: %v\n\n", sendEmailResult)

	return "user signup success", nil
}
