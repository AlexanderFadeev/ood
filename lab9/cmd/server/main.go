package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"mime"
	"net/http"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.InfoLevel)

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".mjs", "application/javascript")
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	server := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	router.Use(logMiddleware)

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}

func logMiddleware(handler http.Handler) http.Handler {
	return &logMiddlewareHandler{handler}
}

type logMiddlewareHandler struct {
	handler http.Handler
}

func (lmh *logMiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"URI": r.RequestURI,
	}).Info()

	lmh.handler.ServeHTTP(w, r)
}
