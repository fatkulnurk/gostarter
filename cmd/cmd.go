package cmd

import (
	"flag"
	"fmt"
	"magicauth/cmd/http"
	"magicauth/cmd/worker"
	"magicauth/config"
	"os"
)

func ServeApp(svc string, cfg *config.Config) {
	switch svc {
	case "http":
		fmt.Println("Running in HTTP server mode...")
		http.Serve(cfg)
	case "worker":
		fmt.Println("Running in background worker mode...")
		worker.Serve(cfg)
	default:
		_, err := fmt.Fprintf(os.Stderr, "Error: invalid --app value: %s\n", svc)
		if err != nil {
			return
		}
		flag.Usage()
		os.Exit(1)
	}
}
