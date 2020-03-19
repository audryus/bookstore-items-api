package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/logger"
)

var (
	//ElasticSearch client
	ElasticSearch elasticSearchInterface = &elasticSearch{}
)

//Init connection to ES
func Init() {
	logger := logger.Get()
	client, err := elastic.NewClient(
		elastic.SetURL("http://local.audryus.com:9200"),
		// elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)

	if err != nil {
		panic(err)
	}
	ElasticSearch.setClient(client)
}

type elasticSearchInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type elasticSearch struct {
	client *elastic.Client
}

func (es *elasticSearch) setClient(cl *elastic.Client) {
	es.client = cl
}

//Index document in ES
func (es *elasticSearch) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := es.client.
		Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in %s", index), err)
		return nil, err
	}
	return result, nil
}

//Get document by ID
func (es *elasticSearch) Get(index string, ID string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := es.client.
		Get().
		Index(index).
		Id(ID).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get document %s in %s", ID, index), err)
		return nil, err
	}
	if !result.Found {
		logger.Info(fmt.Sprintf("No document with ID %s", ID))
		return nil, nil
	}

	return result, nil
}

//Search for document(s)
func (es *elasticSearch) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := es.client.Search(index).Query(query).RestTotalHitsAsInt(false).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}
