package items

import (
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/clients"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

const (
	indexItems = "items"
)

//Save the item
func (i *Item) Save() *errors.RestErr {
	result, err := clients.ElasticSearch.Index(indexItems, i)
	if err != nil {
		return errors.InternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}
