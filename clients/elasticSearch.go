package clients

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var (
	//ElasticSearch client
	ElasticSearch elasticSearchInterface = &elasticSearch{}
)

//Init connection to ES
func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://local.audryus.com:9200"),
		// elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
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
	return es.client.
		Index(index).
		BodyJson(doc).
		Do(ctx)
}
