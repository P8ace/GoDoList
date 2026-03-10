package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/P8ace/GoDoList/internal/adapter/env"
	"github.com/P8ace/GoDoList/package/runner"
)

func main() {

	//setup structured logging
	logHandlerOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, logHandlerOptions))
	slog.SetDefault(log)

	//load configuration from env/secrets management
	dsnstring := env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable")

	// check db connectivity
	// TODO: checking DB connectivity

	//Initialize app
	app := application{
		addr: ":9696",
		dsn: dsnstring,
	}

	mux := app.registerControllers()
	server := app.getServer(mux)

	var ctx context.Context = context.Background()
	var runGroup runner.Group = runner.Group{}

	runGroup.Add(runner.ListenInterrupts(ctx))

	runGroup.Add(func() error {
		slog.Info("Starting HTTP Server", "Address", app.addr)
		return server.ListenAndServe()
	}, func(err error) {
		server.Shutdown(ctx)
	})

	shutdownCause := runGroup.Run()
	slog.Error("Service terminated", "Root Cause", shutdownCause)
}
