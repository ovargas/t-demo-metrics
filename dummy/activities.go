package dummy

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
	"time"
)

type Activities struct{}

func (*Activities) ActivityA(ctx context.Context, payload string) (string, error) {
	return fmt.Sprintf("activity-a-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}

func (*Activities) ActivityB(ctx context.Context, payload string) (string, error) {
	return fmt.Sprintf("activity-b-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}

func (*Activities) ActivityC(ctx context.Context, payload string) (string, error) {
	return fmt.Sprintf("activity-c-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}
