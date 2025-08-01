package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/http/middlewares"
	"api/src/http/routers"
	"api/src/storage"
	"api/src/utils"
)

func main() {
	cfg := config.LoadConfig()

	db, err := storage.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	router := routers.RootRouter(db)

	app := middlewares.AddMiddlewares(router, middlewares.Logger)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: app,
	}

	sm := utils.NewShutdownManager()

	sm.Register(func(ctx context.Context) error {
		fmt.Println("🔒 Closing database connection...")
		return db.Close()
	})

	go func() {
		fmt.Printf("🚀 Starting server on: %s\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v\n", err)
		}
	}()

	sm.WaitAndShutdown(server.Shutdown)
}
