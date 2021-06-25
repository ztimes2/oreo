package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":8080"

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
		middleware.SetHeader(headerAccessControlAllowOrigin, "*"),
	)
	r.Post("/signin", handleSignIn)
	r.Post("/refresh", handleRefresh)
	r.Post("/verify", handleVerify)

	if err := http.ListenAndServe(addr, r); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err.Error())
	}
}
