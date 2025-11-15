package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatkulnurk/gostarter/cmd"
	"github.com/fatkulnurk/gostarter/pkg/config"
	"github.com/fatkulnurk/gostarter/pkg/logging"
)

func main() {
	env := os.Getenv("environment")
	cfg := config.New(env)
	logger := logging.NewSlogLogger(nil)
	logging.InitLogging(logger)

	svc := flag.String("svc", "", "specify application mode: http, worker, scheduler")
	flag.Parse()

	if *svc == "" {
		if _, err := fmt.Fprintf(os.Stderr, "Error: --svc flag is required\n"); err != nil {
			os.Exit(2)
		}
		flag.Usage()
		os.Exit(1)
	}

	cmd.ServeApp(*svc, cfg)
}
