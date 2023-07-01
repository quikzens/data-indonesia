package helper

import "api/entity"

func BuildOrderQuery(param entity.OrderQueryParam) string {
	orderQuery := "created_at DESC"

	if param.Default != "" {
		orderQuery = param.Default
	}

	if param.SortBy != "" {
		orderQuery = param.SortBy + " DESC"
		if param.OrderBy == "asc" {
			orderQuery = param.SortBy + " ASC"
		}
	}

	return orderQuery
}
