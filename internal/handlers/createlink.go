package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

// CreateLink godoc
// @Summary      Сокращает ссылку
// @Description  Получает на вход ссылку и отдаёт в ответе сокращённый вариант
// @Accept		 plain
// @Produce      plain
// @Param        data body string true "URL для сокращения" example(https://ya.ru)
// @Success      201 {string} string "Сокращённая ссылка"
// @Failure		 400 {string} _ "Неверный формат запроса"
// @Failure		 500 {string} _ "Ошибка сервера"
// @Router       / [post] .
func (s *Server) CreateLink(writer http.ResponseWriter, request *http.Request) {
	var resp []byte
	var respStatus int

	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}
	userSession, err := request.Cookie(CookieName)
	if err != nil {
		http.Error(writer, UserUnauthorizedErrTxt, http.StatusBadRequest)
		return
	}

	key := s.Generator.CreateKey()

	err = s.Storage.Add(key, string(body), userSession.Value)
	if err != nil {
		logger.Log.Error(err)
		if errors.Is(err, storage.ErrConflict) {
			shortURL, err := s.Storage.GetByOriginURL(string(body))
			if err != nil {
				http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
				return
			}
			respStatus = http.StatusConflict
			resp = []byte(buildShortURL(s.Config.ResolveHost, shortURL))
		} else {
			http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
			return
		}
	} else {
		respStatus = http.StatusCreated
		resp = []byte(buildShortURL(s.Config.ResolveHost, key))
	}
	writer.WriteHeader(respStatus)
	_, err = writer.Write(resp)
	if err != nil {
		logger.Log.Error(err)
	}
}

// CreateLinkJSON godoc
// @Summary      Сокращает ссылку
// @Description  Получает на вход ссылку и отдаёт в ответе сокращённый вариант
// @Accept       json
// @Produce      json
// @Param 		 data body models.RequestCreateLink true "URL для сокращения"
// @Success      201 {object} models.ResponseCreateLink
// @Failure		 400 {object} object{} "Неверный формат запроса"
// @Failure		 500 {object} object{} "Ошибка сервера"
// @Router       /api/shorten [post] .
func (s *Server) CreateLinkJSON(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLink
	var resp models.ResponseCreateLink
	var respStatus int

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}
	url := body.URL

	userSession, err := request.Cookie(CookieName)
	if err != nil {
		http.Error(writer, UserUnauthorizedErrTxt, http.StatusBadRequest)
		return
	}

	if len(url) == 0 {
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}

	key := s.Generator.CreateKey()
	err = s.Storage.Add(key, url, userSession.Value)
	if err != nil {
		if errors.Is(err, storage.ErrConflict) {
			shortURL, err := s.Storage.GetByOriginURL(url)
			if err != nil {
				http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
				return
			}
			respStatus = http.StatusConflict
			resp = models.ResponseCreateLink{
				Result: buildShortURL(s.Config.ResolveHost, shortURL),
			}
		} else {
			http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
			return
		}
	} else {
		resp = models.ResponseCreateLink{
			Result: buildShortURL(s.Config.ResolveHost, key),
		}
		respStatus = http.StatusCreated
	}

	writer.Header().Set(ContentTypeHeader, ContentTypeJSON)
	writer.WriteHeader(respStatus)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
}
