/*
Run app
*/
package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"testWallet/internal/config"
	"testWallet/internal/database/postgres"
	"testWallet/internal/logger"
	// "testWallet/internal/models"

	// "testWallet/internal/services"
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
	repo:= postgres.NewPGConnection(cfg)
	defer repo.CloseDB()
	
	// u := models.User{
	// 	Name: "Stas5",
	// }
	repo.Update("1", "13")
	repo.Update("2", "13") 
	// fmt.Printf("%+v", a)
	// a := repo.Create(u.Name) 
	// fmt.Printf("%+v", a)
	// if err := repo.Delete(u.Name); err!= nil{
	// 	fmt.Println(err)
	// }

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

