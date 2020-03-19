package queries

type EsQuery struct {
	Equals []FieldValue
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
