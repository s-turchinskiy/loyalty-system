package handlers

import (
	"errors"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func (h *Handler) NewOrder(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Type"), ContentTypeTextPlain) {

		err := fmt.Errorf("Content-Type != %s", ContentTypeTextPlain)
		errorGettingData(w, err)
		return
	}

	defer r.Body.Close()
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		errorGettingData(w, err)
		return
	}

	context := r.Context()
	login := context.Value(userLogin).(string)
	orderID := string(responseData)
	err = h.Service.NewOrder(context, login, orderID)

	switch {

	case errors.Is(err, servicecommon.ErrNoLuhnValidate):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	case errors.Is(err, servicecommon.ErrOrderNumberAlreadyUploadedByThisUser):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusOK)
		return
	case errors.Is(err, servicecommon.ErrOrderNumberAlreadyUploadedByAnotherUser):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusConflict)
		return
	case err != nil:
		internalError(w, err)
		return

	}

	w.WriteHeader(http.StatusAccepted)

}
