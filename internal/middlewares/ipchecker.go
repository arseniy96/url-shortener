package middlewares

import (
	"net"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
)

const (
	InternalBackendErrTxt  = "Internal Backend Error"
	UserUnauthorizedErrTxt = "User unauthorized"
)

func IPCheckerMiddleware(trustedSubnet string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path != "/api/internal/stats" {
				h.ServeHTTP(writer, request)
				return
			}

			if trustedSubnet == "" {
				logger.Log.Error("trusted subnet is empty")
				http.Error(writer, UserUnauthorizedErrTxt, http.StatusForbidden)
				return
			}
			_, ipNet, err := net.ParseCIDR(trustedSubnet)
			if err != nil {
				logger.Log.Errorf("parse subnet error: %v", err)
				http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
				return
			}

			xRealIP := request.Header.Get("X-Real-IP")
			if xRealIP == "" {
				logger.Log.Error("X-Real-IP is empty")
				http.Error(writer, UserUnauthorizedErrTxt, http.StatusForbidden)
				return
			}
			realIP := net.ParseIP(xRealIP)
			if !ipNet.Contains(realIP) {
				logger.Log.Error("IP is not from trusted subnet")
				http.Error(writer, UserUnauthorizedErrTxt, http.StatusForbidden)
				return
			}
			h.ServeHTTP(writer, request)
		})
	}
}
