package responseHelper

import (
	"database/sql"
	"net/http"

	"github.com/zeintkp/go-rest/domain"
)

//BuildResponse is used to send response
func BuildResponse(statusCode int, data, pagination interface{}, err error) (int, interface{}) {

	if err != nil {

		if err == sql.ErrNoRows {
			statusCode = http.StatusNotFound
			data = nil
		} else if statusCode < 400 {
			statusCode = http.StatusUnprocessableEntity
			data = domain.ErrorVM{
				Message: err.Error(),
			}
		} else {
			data = domain.ErrorVM{
				Message: err.Error(),
			}
		}
	}

	return statusCode, domain.ResponseVM{
		Code:       statusCode,
		Status:     http.StatusText(statusCode),
		Data:       data,
		Pagination: pagination,
	}
}
