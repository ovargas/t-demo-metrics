package dummy

import (
	"fmt"
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
	})

	if err := workflow.ExecuteActivity(ctx, a.ActivityA, payload).Get(ctx, &aResult); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, a.ActivityB, payload).Get(ctx, &bResult); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, a.ActivityC, fmt.Sprintf("%s|%s", aResult, bResult)).Get(ctx, &bResult); err != nil {
		return err
	}

	return nil
}
