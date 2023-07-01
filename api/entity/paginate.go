package entity

type PaginateResult[T any] struct {
	Data   []T
	Total  int
	Limit  int
	Offset int
}

type BasePaginateParam struct {
	Keywords string
	SortBy   string
	OrderBy  string
	Limit    int
	Offset   int
}
