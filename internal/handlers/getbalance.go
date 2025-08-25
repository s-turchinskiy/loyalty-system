package handlers

import (
	"github.com/mailru/easyjson"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {

	context := r.Context()
	result, err := h.Service.GetBalance(context, "")
	if err != nil {
		internalError(w, err)
		return
	}

	rawBytes, err := easyjson.Marshal(result)
	if err != nil {
		logger.Log.Info("error encoding response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeApplicationJson)
	w.Write(rawBytes)

}
