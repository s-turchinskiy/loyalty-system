package handlers

import (
	"encoding/json"
	"errors"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) NewWithdraw(w http.ResponseWriter, r *http.Request) {

	var req models.NewWithdraw
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		loggingCannotDecodeRequestJSONBody(r.Body)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	context := r.Context()
	err := h.Service.NewWithdraw(context, "", req)

	switch {
	case errors.Is(err, common.ErrorNotEnoughBalance):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusPaymentRequired)

	case common.IsErrorDuplicateKeyValue(err):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusUnprocessableEntity)

	case err != nil:
		internalError(w, err)
		return
	}

}

func loggingCannotDecodeRequestJSONBody(body io.ReadCloser) {

	bodyByte, err := io.ReadAll(body)
	if err != nil {
		logger.Log.Info("cannot decode request JSON body", zap.Error(common.WrapError(err)))
		return
	}

	logger.Log.Info("cannot decode request JSON body", zap.Error(common.WrapError(err)), zap.String("body", string(bodyByte)))

}
