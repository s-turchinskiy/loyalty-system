package handlers

import (
	"context"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"net/http"
	"strings"
)

func (h *Handler) AuthorizationMiddleware(next http.Handler) http.Handler {
	authorizationFn := func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			logger.Log.Debugw(common.WrapError(fmt.Errorf("authorization header is empty")).Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		token := strings.Split(authorizationHeader, " ")[1]
		success, login, err := h.tokenVerification(ctx, token)
		if err != nil {
			logger.Log.Debugw(common.WrapError(err).Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !success {
			logger.Log.Debugw(common.WrapError(fmt.Errorf("authorization failed")).Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, "login", login)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(authorizationFn)
}
