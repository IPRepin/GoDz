package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.WithFields(logrus.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"statusCode": wrapper.StatusCode,
			"duration":   duration.String(),
		}).Info("Request handled")
	})
}
