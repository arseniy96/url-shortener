package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

// DeleteUserUrls godoc
// @Summary      Удаляет ссылки пользователя
// @Description  Получает на вход массив ссылок и удаляет их
// @Accept       json
// @Produce      plain
// @Param 		 data body models.RequestDeleteUserURLS true "URL для удаления"
// @Success      202 {string} string{ok}
// @Failure		 400 {object} object{} "Неверный формат запроса"
// @Failure		 401 {object} object{} "Unauthorized"
// @Router       /api/user/urls [delete] .
func (s *Server) DeleteUserUrls(writer http.ResponseWriter, request *http.Request) {
	var urls models.RequestDeleteUserURLS

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&urls); err != nil {
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}

	userCookie, err := request.Cookie(CookieName)
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, InvalidCookieErrTxt, http.StatusUnauthorized)
		return
	}

	s.DeleteURLSChan <- storage.DeleteURLMessage{
		ShortURLs:  urls,
		UserCookie: userCookie.Value,
	}

	writer.WriteHeader(http.StatusAccepted)
	_, err = writer.Write([]byte("ok"))
	if err != nil {
		logger.Log.Error(err)
	}
}
