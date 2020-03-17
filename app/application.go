package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/clients"
)

var (
	router = mux.NewRouter()
)

//StartApplication app
func StartApplication() {
	clients.Init()

	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8082",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
