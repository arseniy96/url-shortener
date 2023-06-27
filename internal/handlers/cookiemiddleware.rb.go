package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

const SecretKey = "psqy18SWnnqU9"

func (s *Server) CookieMiddleware(h http.HandlerFunc) http.HandlerFunc {
	// достаём куку и дешифруем её
	// если куки нет, смотрим на роут, если /api/user/urls, то отдаём ошибку
	// если есть кука, проверяем id
	// если куки нет, или id неправильный – генерируем новую куку и сетим её
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		cookie, err := r.Cookie("shortener_session")
		if err != nil {
			newCookie, err := createNewCookie(s.storage)
			if err != nil {
				logger.Log.Error("create cookie error", zap.Error(err))
				http.Error(w, "Internal Backend Error", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:  "shortener_session",
				Value: newCookie,
			})
			r.AddCookie(&http.Cookie{
				Name:  "shortener_session",
				Value: newCookie,
			})
		} else if !cookieValid(cookie.Value, s.storage) {
			if path == "/api/user/urls" {
				logger.Log.Error("invalid cookie", zap.Error(err))
				http.Error(w, "invalid cookie", http.StatusUnauthorized)
				return
			}

			newCookie, err := createNewCookie(s.storage)
			if err != nil {
				logger.Log.Error("create cookie error", zap.Error(err))
				http.Error(w, "Internal Backend Error", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:  "shortener_session",
				Value: newCookie,
			})
			r.AddCookie(&http.Cookie{
				Name:  "shortener_session",
				Value: newCookie,
			})
		}
		h.ServeHTTP(w, r)
	}
}

func createNewCookie(rep Repository) (cookie string, err error) {
	// создаём нового пользователя в базе и берём его id
	// генерируем новую куку
	// сохраняем куку в пользователя
	if rep.GetMode() == storage.DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		user, err := rep.CreateUser(ctx)
		if err != nil {
			return "", err
		}
		userID := user.UserID

		cookie, err = BuildJWTString(userID)
		if err != nil {
			return cookie, err
		}
		err = rep.UpdateUser(ctx, userID, cookie)
		if err != nil {
			return cookie, err
		}
	} else {
		logger.Log.Infow("not db mode")
		userID := 1
		cookie, err = BuildJWTString(userID)
		if err != nil {
			return cookie, err
		}
	}

	return cookie, nil
}

func BuildJWTString(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func cookieValid(cookie string, rep Repository) bool {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		logger.Log.Error("parse jwt error", zap.Error(err))
		return false
	}

	if !token.Valid {
		logger.Log.Error("Token is not valid")
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	userID := claims.UserID
	_, err = rep.FindUserByID(ctx, userID)
	return err == nil
}
