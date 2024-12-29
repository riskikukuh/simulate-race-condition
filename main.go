package main

import (
	"log/slog"
	"simulation-race-condition/config"
	"simulation-race-condition/database"
	"simulation-race-condition/job"
)

func main() {
	slog.Info("Start Application")

	defer func() {
		slog.Info("Stopping Application")
	}()

	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	wrapDB, err := database.InitDatabase(env)
	if err != nil {
		panic(err)
	}

	job.StartBalanceJob(env, wrapDB)
}
