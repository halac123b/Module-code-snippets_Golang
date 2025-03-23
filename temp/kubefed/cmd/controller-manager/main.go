package main

import (
	"fmt"
	"os"

	"k8s.io/component-base/logs"
)

// Controller-manager main.
func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	stopChan := genericapiserver.SetupSignalHandler()

	if err := app.NewControllerManagerCommand(stopChan).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1) //nolint:gocritic
	}
}
