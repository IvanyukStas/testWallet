/*
Run app
*/
package app

import (
	"log/slog"
	"testWallet/internal/config"
	"testWallet/internal/logger"
)

//Run app
func Run(){
	
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log = log.With(slog.String("ENV", cfg.Env))
	log.Info("initializing server", slog.String("adress", cfg.Host))
	log.Debug("start the program")
	
	
}