package dummy

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
	"math/rand"
	"time"
)

type Activities struct{}

func (*Activities) ActivityA(ctx context.Context, payload string) (string, error) {
	simProcess(10)

	if payload == "fail" {
		return "", ErrUpsIAteABug
	}

	return fmt.Sprintf("activity-a-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}

func (*Activities) ActivityB(ctx context.Context, payload string) (string, error) {
	simProcess(10)
	n := rand.Int31n(100)

	if n > 90 {
		return "", fmt.Errorf("any error in activity a")
	}

	return fmt.Sprintf("activity-b-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}

func (*Activities) ActivityC(ctx context.Context, payload string) (string, error) {
	simProcess(15)
	n := rand.Int31n(100)
	if n > 40 {
		return "", fmt.Errorf("any error in activity b")
	}

	return fmt.Sprintf("activity-c-%s-%s-%d", payload, activity.GetInfo(ctx).WorkflowExecution.RunID, time.Now().UnixMilli()), nil
}

func simProcess(n int32) {
	t := rand.Int31n(n)
	time.Sleep(time.Second * time.Duration(t))
}
