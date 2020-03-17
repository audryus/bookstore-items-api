package services

import (
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/items"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	//ItemService instance
	ItemService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}
type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, errors.NotImpemented()
}

func (s *itemService) Get(ID string) (*items.Item, *errors.RestErr) {
	return nil, errors.NotImpemented()
}
