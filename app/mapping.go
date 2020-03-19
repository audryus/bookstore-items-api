package app

import (
	"net/http"

	controller "gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controller.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controller.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controller.ItemController.Search).Methods(http.MethodPost)
}
