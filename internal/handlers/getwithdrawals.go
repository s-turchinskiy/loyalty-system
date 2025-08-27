package handlers

import (
	"github.com/mailru/easyjson"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetWithdrawals(w http.ResponseWriter, r *http.Request) {

	context := r.Context()
	login := context.Value("login").(string)
	withdrawals, err := h.Service.GetWithdrawals(context, login)
	if err != nil {
		internalError(w, err)
		return
	}

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	rawBytes, err := easyjson.Marshal(withdrawals)
	if err != nil {
		logger.Log.Info("error encoding response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeApplicationJSON)
	w.Write(rawBytes)

}
