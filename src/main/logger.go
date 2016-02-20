package main
import (
	"net/http"
	"github.com/Sirupsen/logrus"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.WithFields(logrus.Fields{
			"method": r.Method,
			"requestURI": r.RequestURI,
			"name": name,
			"time": time.Since(start),
		}).Info("HTTP request")
	})
}