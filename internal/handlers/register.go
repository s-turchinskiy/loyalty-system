package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/authservice"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Type"), ContentTypeApplicationJSON) {

		err := fmt.Errorf("Content-Type != %s", ContentTypeApplicationJSON)
		errorGettingData(w, err)
		return
	}

	var req models.NewUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		loggingCannotDecodeRequestJSONBody(r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	context := r.Context()
	hashPassword, err := h.Service.Register(context, req)

	switch {
	case errors.Is(err, servicecommon.ErrUserAlreadyExist):
		logger.Log.Info(zap.Error(err))
		w.WriteHeader(http.StatusUnauthorized)

	case err != nil:
		internalError(w, err)
		return
	}

	token, err := authservice.BuildJWTString(req.Login, hashPassword, h.tokenExp)
	w.Header().Set("Authorization", "Bearer "+token)
}
