package app

import (
	"net/http"

	controller "gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controller.ItemController.Create).Methods(http.MethodPost)
}
