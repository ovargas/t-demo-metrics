package main

import (
	"context"
	"flag"
	"fmt"
	"go.temporal.io/sdk/client"
	"log"
	"t-demo-metric/dummy"
)

func main() {

	payload := flag.String("payload", "", "a string")
	flag.Parse()

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}

	wr, err := c.ExecuteWorkflow(context.TODO(),
		client.StartWorkflowOptions{
			TaskQueue: "dummy",
		},
		dummy.Dummy, *payload)

	fmt.Printf("%s", wr.GetID())
}
