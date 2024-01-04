package main

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
	"test-idp/cmd"
	"test-idp/internal/config"
)

func main() {
	cnf, err := config.Parse()
	if err != nil {
		panic(err)
	}

	stopSignals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	ctx, _ := signal.NotifyContext(context.Background(), stopSignals...)

	log := zerolog.New(os.Stdout)

	var exitCode int
	if err = cmd.Run(ctx, cnf); err != nil {
		log.Err(err).Send()
		exitCode = 1
	}

	os.Exit(exitCode)
}
