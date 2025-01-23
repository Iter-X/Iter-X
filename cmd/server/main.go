package main

import (
	"context"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/log"
	"github.com/iter-x/iter-x/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cfgPath string
	var rootCmd = &cobra.Command{
		Use:   "iter-x",
		Short: "CLI for managing Iter X Server",
		Long:  `The Iter X Server CLI provides a command-line interface for managing and interacting with the Iter X Server service.`,
		Run:   func(cmd *cobra.Command, args []string) { run(cfgPath) },
	}

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "conf", "c", "./config", "Path to the configuration files")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cfgPath string) {
	cfg, err := conf.Load(cfgPath)
	if err != nil {
		panic(err)
	}

	logger, err := log.New(cfg.Environment == conf.Environment_DEV, cfg.Log)
	if err != nil {
		panic(err)
	}

	logger.Info("Iter X Server is initializing...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Info("Iter X Server initialized successfully, starting the server...")

	app, cleanup, err := wireApp(cfg.I18N, cfg.Server.Grpc, cfg.Server.Http, cfg.Data, cfg.Auth, cfg.Agent, logger)
	if err != nil {
		logger.Error("Failed to init the app", err)
		os.Exit(1)
	}

	startServer("gRPC", app.GrpcServer.Start, logger)
	startServer("HTTP", func() error {
		return app.HttpServer.Start(ctx)
	}, logger)

	logger.Info("Iter X Server started successfully")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-c:
			logger.Info("Received signal to stop the server")
			cancel()
		case <-ctx.Done():
			app.GrpcServer.Stop()
			cleanup()
			logger.Info("Iter X Server stopped")
			return
		}
	}
}

type App struct {
	GrpcServer *server.GRPCServer
	HttpServer *server.HTTPServer
}

func newApp(grpcServer *server.GRPCServer, httpServer *server.HTTPServer) *App {
	return &App{
		GrpcServer: grpcServer,
		HttpServer: httpServer,
	}
}

func startServer(name string, startFunc func() error, logger *zap.SugaredLogger) {
	go func() {
		if err := startFunc(); err != nil {
			logger.Errorw(name+" server encountered an error", "error", err)
			os.Exit(1)
		}
	}()
}
