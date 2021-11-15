package domain

type ResponseVM struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

//ErrorVM struct for error response
type ErrorVM struct {
	Message interface{} `json:"message"`
}

//PaginationVM struct for pagination
type PaginationVM struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}
