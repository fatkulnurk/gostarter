package main

import (
	"flag"
	"fmt"
	"magicauth/cmd"
	"magicauth/config"
	"os"
)

func main() {
	cfg := config.New(os.Getenv("environment"))

	svc := flag.String("svc", "", "specify application mode: http or worker")
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
