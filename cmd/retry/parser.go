package main

import (
	"os"
	"time"

	"github.com/kamilsk/retry/cmd/retry/flag"
	"github.com/kamilsk/retry/strategy"
)

func parse() (time.Duration, []string, []strategy.Strategy) {
	cl := flag.NewSet("retry")
	cl.Usage = usage
	for name, cfg := range compliance {
		switch cursor := cfg.cursor.(type) {
		case *string:
			cl.StringVar(cursor, name, "", cfg.usage)
		case *bool:
			cl.BoolVar(cursor, name, false, cfg.usage)
		}

	}
	cl.StringVar(&Timeout, "timeout", Timeout, "value which supported by time.ParseDuration")
	if err := cl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	timeout, err := time.ParseDuration(Timeout)
	if err != nil {
		panic(err)
	}

	strategies, err := handle(cl.Sequence())
	if err != nil {
		panic(err)
	}

	args := cl.Args()
	if len(args) == 0 {
		panic("please provide a command to retry")
	}

	return timeout, cl.Args(), strategies
}