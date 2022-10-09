package dummy

import (
	"fmt"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func Dummy(ctx workflow.Context, payload string) error {
	var a Activities

	var (
		aResult string
		bResult string
	)

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			NonRetryableErrorTypes: []string{
				"Error",
			},
		},
	})

	if err := workflow.ExecuteActivity(ctx, a.ActivityA, payload).Get(ctx, &aResult); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, a.ActivityB, payload).Get(ctx, &bResult); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, a.ActivityC, fmt.Sprintf("%s|%s", aResult, bResult)).Get(ctx, &bResult); err != nil {
		return err
	} else {
		if payload == "panic" {
			// Make the workflow panic by invalid memory address or nil pointer dereference
			workflow.GetLogger(ctx).Info("making the workflow panic", err.Error())
		}
	}

	return nil
}
