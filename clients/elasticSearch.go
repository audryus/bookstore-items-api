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
}

type elasticSearch struct {
	client *elastic.Client
}

func (es *elasticSearch) setClient(cl *elastic.Client) {
	es.client = cl
}

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
