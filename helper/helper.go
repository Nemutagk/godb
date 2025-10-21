package helper

import (
	"fmt"
	"strings"

	"github.com/Nemutagk/godb/v2/definitions/models"
)

func PrepareFilters(filters models.GroupFilter, counter int) (string, []any, int) {
	var queryBuilder strings.Builder

	if counter <= 0 {
		counter = 1
	}

	vals := []any{}
	for _, tmpFilter := range filters.Filters {

		var currentParts strings.Builder
		var currentVals []any

		if filter, ok := tmpFilter.(models.Filter); ok {
			comparator := "="
			if filter.Comparator != nil {
				comparator = *filter.Comparator
			}

			if comparator != models.FilterSqlComparatorIsNull && comparator != models.FilterSqlComparatorIsNotNull {
				currentParts.WriteString(fmt.Sprintf("%s %s $%d", filter.Key, comparator, counter))
				currentVals = append(currentVals, filter.Value)

				counter++
			} else {
				currentParts.WriteString(fmt.Sprintf("%s %s", filter.Key, comparator))
			}
		} else if groupFilter, ok := tmpFilter.(models.GroupFilter); ok {
			subQuery, subVals, newCounter := PrepareFilters(groupFilter, counter)
			counter = newCounter

			if subQuery != "" {
				currentParts.WriteString(fmt.Sprintf("(%s)", subQuery))
				currentVals = append(currentVals, subVals...)
			}
		}

		if currentParts.Len() > 0 {
			if queryBuilder.Len() > 0 {
				groupOperator := "AND"
				if filters.Operator != "" {
					groupOperator = filters.Operator
				}
				queryBuilder.WriteString(fmt.Sprintf(" %s ", groupOperator))
			}

			queryBuilder.WriteString(currentParts.String())
			vals = append(vals, currentVals...)
		}
	}

	return queryBuilder.String(), vals, counter
}

func PrepareSoftDelete(softDelete *string, filters models.GroupFilter) models.GroupFilter {
	if softDelete == nil || *softDelete == "" {
		return filters
	}

	newGroup := models.GroupFilter{
		Operator: "OR",
		Filters:  []any{},
	}

	// opAnd := models.FilterOperatorAnd
	or := "OR"
	isNull := models.FilterSqlComparatorIsNull

	newGroup.Filters = append(newGroup.Filters, models.Filter{
		Key:        *softDelete,
		Comparator: &isNull,
		Value:      nil,
		Operator:   &or,
	})

	if len(filters.Filters) > 0 {
		filters.Filters = append(filters.Filters, newGroup)
	} else {
		filters = newGroup
	}

	return filters
}
