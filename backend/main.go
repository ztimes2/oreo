package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":8080"

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Info("server is listening on " + addr)
	
	if err := http.ListenAndServe(addr, newRouter()); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err.Error())
	}
}
