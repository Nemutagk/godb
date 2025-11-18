package models

const (
	FilterSqlOperatorAnd = "AND"
	FilterSqlOperatorOr  = "OR"

	FilterSqlComparatorEqual        = "="
	FilterSqlComparatorNotEqual     = "!="
	FilterSqlComparatorGreaterThan  = ">"
	FilterSqlComparatorLessThan     = "<"
	FilterSqlComparatorGreaterEqual = ">="
	FilterSqlComparatorLessEqual    = "<="
	FilterSqlComparatorLike         = "LIKE"
	FilterSqlComparatorIsNull       = "IS NULL"
	FilterSqlComparatorIsNotNull    = "IS NOT NULL"
)

type Filter struct {
	Key        string  `json:"key"`
	Value      any     `json:"value"`
	Operator   *string `json:"operator"`
	Comparator *string `json:"comparator"`
}

type FilterMultipleValue struct {
	Key        string  `json:"key"`
	Values     []any   `json:"values"`
	Operator   *string `json:"operator"`
	Comparator *string `json:"comparator"`
}

type GroupFilter struct {
	Filters  []any  `json:"filters"`
	Operator string `json:"operator"`
}

type Options struct {
	Columns          *[]string
	OrderColumn      string
	OrderDir         string
	Limit            int64
	Offset           int64
	Relations        []string
	PrimaryKey       *string
	InsertPrimaryKey *bool
	TimestampsFields *bool
}
