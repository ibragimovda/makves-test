package app

import (
	"golang.org/x/exp/slog"
	"makves/internal/config"
	"makves/internal/repository/pgrepo"
	"makves/internal/service"
	"makves/internal/storage/pg"
	"makves/internal/transport/rest/handler"
	"net/http"
	"os"
)

func Run() {
	cfg := config.MustLoad()
	//logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	//storage
	db := pg.New()
	repo := pgrepo.New(db)
	itemService := service.NewService(repo)
	mux := handler.NewHandler(itemService, log)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      mux.InitRoutes(),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("Ошибка при запуске сервера")
	}

	log.Error("Сервер остановлен")
}
