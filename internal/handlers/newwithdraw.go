package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func (h *Handler) NewWithdraw(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Type"), ContentTypeApplicationJSON) {

		err := fmt.Errorf("Content-Type != %s", ContentTypeApplicationJSON)
		errorGettingData(w, err)
		return
	}

	var req models.NewWithdraw
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		loggingCannotDecodeRequestJSONBody(r.Body)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	context := r.Context()
	login := context.Value("login").(string)
	err := h.Service.NewWithdraw(context, login, req)

	switch {
	case errors.Is(err, servicecommon.ErrorNotEnoughBalance):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusPaymentRequired)

	case servicecommon.IsErrorDuplicateKeyValue(err):
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
