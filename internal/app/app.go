/*
Run app
*/
package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"testWallet/internal/config"
	"testWallet/internal/logger"
	"testWallet/internal/services"
	// mwLogger "testWallet/internal/transport/rest/middleware/logger"
	hh "testWallet/internal/transport/rest/handler"

	// "github.com/justinas/alice"
)

// Run app
func Run() {

	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log = log.With(slog.String("ENV", cfg.Env))
	log.Info("initializing server", slog.String("adress", cfg.Host))
	log.Debug("start the program")

	// middlewareChain := alice.New(mwLogger.New(log))
	repo, _ := services.NewPostgresRepository()

	h := &hh.WalletHandler{
		Repo: repo,
		Logger: log,
	}

	
	mux := http.NewServeMux()
	// finalHandler := http.HandlerFunc(final)
	// mux.Handle("/", middlewareChain.Then(finalHandler))

	mux.HandleFunc("/", h.Post)

	err := http.ListenAndServe(":8080", mux)
	fmt.Println(err)

	

}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

