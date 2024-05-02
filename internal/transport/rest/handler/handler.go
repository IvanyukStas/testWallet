package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"testWallet/internal/models"
)

type WalletHandler struct {
	Repo   models.Cruder
	Logger *slog.Logger
}

// HTTP POSTapi wallet/useradd
func (wh *WalletHandler) Post(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		wh.Logger.Error(err.Error(), slog.String("url", r.URL.String()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = wh.Repo.Create(user.Name); err != nil {
		wh.Logger.Info(err.Error(), slog.String("url", r.URL.String()))
		if err == models.ErrorUserExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wh.Logger.Info("Create new user and wallet", slog.String("url", r.URL.String()))
	w.WriteHeader(http.StatusCreated)
}
