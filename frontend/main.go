package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":8081"

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Info("server is listening on " + addr)

	r := chi.NewRouter()
	r.Use(
		middleware.RequestLogger(&middleware.DefaultLogFormatter{
			Logger: logrus.StandardLogger(),
		}),
		middleware.Recoverer,
	)
	r.Mount("/", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(addr, r); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err.Error())
	}
}
