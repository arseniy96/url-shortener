package handlers

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
)

func (s *Server) Stats(writer http.ResponseWriter, request *http.Request) {
	trustedSubnet := s.Config.TrustedSubnet
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

	ctx, cancelCtx := context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()
	urlsCount, err := s.Storage.GetURLsCount(ctx)
	if err != nil {
		logger.Log.Errorf("get urls count error: %v", err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	ctx, cancelCtx = context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()
	usersCount, err := s.Storage.GetUsersCount(ctx)
	if err != nil {
		logger.Log.Errorf("get users count error: %v", err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	resp := models.ResponseStats{
		URLs:  urlsCount,
		Users: usersCount,
	}
	writer.Header().Set(ContentTypeHeader, ContentTypeJSON)
	writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
}
