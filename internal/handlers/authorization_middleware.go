package handlers

import (
	"context"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type login string

const (
	userLogin login = "login"
)

func (h *Handler) AuthorizationMiddleware(next http.Handler) http.Handler {
	authorizationFn := func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/api/user/register" || r.RequestURI == "/api/user/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {

			var headers []string
			for name, values := range r.Header {
				for _, value := range values {
					headers = append(headers, fmt.Sprintf("%s:%s", name, value))
				}
			}

			err := fmt.Errorf("authorization header is empty, url: %s, headers: %s",
				r.RequestURI,
				strings.Join(headers, "\n"))

			logger.Log.Debug("authorization header is empty", zap.Error(common.WrapError(err)))
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

		ctx = context.WithValue(ctx, userLogin, login)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(authorizationFn)
}
