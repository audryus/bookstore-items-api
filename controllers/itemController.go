package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/items"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/queries"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/services"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/utils"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-go/oauth"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	//ItemController item controller
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemController struct{}

//Create item
func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		utils.HTTP.ResponseError(w, err)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		utils.HTTP.ResponseError(w, errors.UnautorizedError())
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HTTP.ResponseError(w, errors.BadRequestError("invalid request body", err))
		return
	}
	defer r.Body.Close()
	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		utils.HTTP.ResponseError(w, errors.BadRequestError("invalid item json body", err))
		return
	}
	itemRequest.Seller = sellerID
	result, createErr := services.ItemService.Create(itemRequest)

	if createErr != nil {
		utils.HTTP.ResponseError(w, createErr)
		return
	}
	utils.HTTP.ResponseJSON(w, http.StatusCreated, result)
}

//Get item by id
func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Get(itemID)
	if err != nil {
		utils.HTTP.ResponseError(w, err)
		return
	}

	utils.HTTP.ResponseJSON(w, http.StatusOK, item)
}

//Get item by id
func (c *itemController) Search(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HTTP.ResponseError(w, errors.BadRequestError("invalid json body", err))
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery

	if err := json.Unmarshal(requestBody, &query); err != nil {
		utils.HTTP.ResponseError(w, errors.BadRequestError("invalid query json body", err))
		return
	}

	items, err := services.ItemService.Search(query)
	if err != nil {
		utils.HTTP.ResponseError(w, errors.BadRequestError("invalid search criteria", err))
		return
	}
	utils.HTTP.ResponseJSON(w, http.StatusOK, items)
}
