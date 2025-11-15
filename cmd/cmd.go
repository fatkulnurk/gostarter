package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatkulnurk/gostarter/cmd/http"
	"github.com/fatkulnurk/gostarter/cmd/scheduler"
	"github.com/fatkulnurk/gostarter/cmd/worker"
	"github.com/fatkulnurk/gostarter/pkg/config"
)

func ServeApp(svc string, cfg *config.Config) {
	switch svc {
	case "http":
		fmt.Println("Running in HTTP server mode...")
		http.Serve(cfg)
	case "worker":
		fmt.Println("Running in background worker mode...")
		worker.Serve(cfg)
	case "scheduler":
		fmt.Println("Running in scheduler mode...")
		scheduler.Serve(cfg)
	default:
		_, err := fmt.Fprintf(os.Stderr, "Error: invalid --svc value: %s\n", svc)
		if err != nil {
			return
		}
		flag.Usage()
		os.Exit(1)
	}
}
