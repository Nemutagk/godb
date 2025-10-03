package models

const (
	FilterOperatorAnd = "AND"
	FilterOperatorOr  = "OR"

	FilterComparatorEqual        = "="
	FilterComparatorNotEqual     = "!="
	FilterComparatorGreaterThan  = ">"
	FilterComparatorLessThan     = "<"
	FilterComparatorGreaterEqual = ">="
	FilterComparatorLessEqual    = "<="
	FilterComparatorLike         = "LIKE"
	FilterComparatorIsNull       = "IS NULL"
	FilterComparatorIsNotNull    = "IS NOT NULL"
)

type Filter struct {
	Key        string  `json:"key"`
	Value      any     `json:"value"`
	Operator   *string `json:"operator"`
	Comparator *string `json:"comparator"`
}

type GroupFilter struct {
	Filters  []any  `json:"filters"`
	Operator string `json:"operator"`
}

type Options struct {
	Columns     *[]string
	OrderColumn string
	OrderDir    string
	Limit       int64
	Offset      int64
}
