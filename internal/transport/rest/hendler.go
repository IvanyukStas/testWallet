package rest

import (
	"log/slog"
	"net/http"
	"testWallet/internal/models"
)

type WalletHandler struct {
	Repo   models.Cruder
	Logger *slog.Logger
}

// HTTP POSTapi wallet/useradd
func (wh WalletHandler) Post(w http.ResponseWriter, r http.Request) {
	
}
