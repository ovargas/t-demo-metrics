package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"log"
	"t-demo-metric/dummy"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}

	wr, err := c.ExecuteWorkflow(context.TODO(),
		client.StartWorkflowOptions{
			TaskQueue: "dummy",
		},
		dummy.Dummy, "")

	fmt.Printf("%s", wr.GetID())
}
