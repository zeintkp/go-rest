package pagination

import "github.com/zeintkp/go-rest/domain"

const (
	defaultLimit = 10
	maxLimit     = 100
	defaultOrder = "id"
	defaultSort  = "asc"
)

//SetPaginationParameter is used to set pagination parameter for query data
func SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	if order == "" {
		order = defaultOrder
	}

	if sort != "desc" {
		sort = defaultSort
	}

	offset := (page - 1) * limit

	return offset, page, limit, order, sort
}

//SetPaginationResponse is used to set pagination response
func SetPaginationResponse(page, limit, total int) domain.PaginationVM {
	var lastPage int

	if total > 0 {
		lastPage = total / limit
		if total%limit != 0 {
			lastPage++
		}
	}

	return domain.PaginationVM{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}
}
