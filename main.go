package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatkulnurk/gostarter/cmd"
	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/pkg/logging"
)

func main() {
	cfg := config.New(os.Getenv("environment"))

	// logging
	logging.InitLogger()

	svc := flag.String("svc", "", "specify application mode: http, worker, scheduler")
	flag.Parse()
	if *svc == "" {
		_, err := fmt.Fprintf(os.Stderr, "Error: --svc flag is required\n")
		if err != nil {
			return
		}
		flag.Usage()
		os.Exit(1)
	}

	cmd.ServeApp(*svc, cfg)
}
