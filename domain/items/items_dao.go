package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/clients"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/items-api/domain/queries"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

const (
	indexItems = "items"
)

//Save the item
func (i *Item) Save() errors.RestErr {
	result, err := clients.ElasticSearch.Index(indexItems, i)
	if err != nil {
		return errors.InternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

//Get the item by ID
func (i *Item) Get() errors.RestErr {
	tmpID := i.ID
	result, err := clients.ElasticSearch.Get(indexItems, tmpID)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.NotFoundError("No item found")
		}
		return errors.InternalServerError(fmt.Sprintf("error whe trying to get ID %s", i.ID),
			errors.New("ES error"))
	}
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("no items found with ID %s", i.ID))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return errors.InternalServerError("error trying to parse database",
			errors.New("ES error"))
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return errors.InternalServerError("error trying to parse response",
			errors.New("database error"))
	}
	i.ID = tmpID
	return nil
}

//Search document with query
func (i *Item) Search(query queries.EsQuery) ([]Item, errors.RestErr) {
	result, err := clients.ElasticSearch.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors.InternalServerError("error when trying to search documents", errors.New("ES error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.InternalServerError("error when trying to parse response", errors.New("ES error"))
		}
		items[index] = item
	}
	if len(items) == 0 {
		return nil, errors.NotFoundError("no items found matching given criteria")
	}
	return items, nil
}
