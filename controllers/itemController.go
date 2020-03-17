package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/items"
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
}

type itemController struct{}

//Create item
func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		utils.HTTP.ResponseError(w, err)
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
	itemRequest.Seller = oauth.GetCallerID(r)
	result, createErr := services.ItemService.Create(itemRequest)

	if createErr != nil {
		utils.HTTP.ResponseError(w, createErr)
		return
	}
	utils.HTTP.ResponseJSON(w, http.StatusCreated, result)
}

//Get item by id
func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}