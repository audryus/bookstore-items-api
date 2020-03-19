package queries

import "github.com/olivere/elastic"

//Build elastic Query
func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	if len(q.Equals) > 0 {

	}
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQueries...)

	return query
}