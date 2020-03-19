package services

import (
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/items"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/queries"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	//ItemService instance
	ItemService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, errors.RestErr)
}
type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(ID string) (*items.Item, errors.RestErr) {
	item := &items.Item{
		ID: ID,
	}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
