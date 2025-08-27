package handlers

import (
	"github.com/mailru/easyjson"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {

	context := r.Context()
	login := context.Value(userLogin).(string)
	orders, err := h.Service.GetOrders(context, login)
	if err != nil {
		internalError(w, err)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	rawBytes, err := easyjson.Marshal(orders)
	if err != nil {
		logger.Log.Info("error encoding response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeApplicationJSON)
	w.Write(rawBytes)

}
