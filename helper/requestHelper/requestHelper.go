package requestHelper

import (
	"errors"
	"go-rest/helper/exception"
	"go-rest/helper/str"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type IRequestHelper interface {
	ValidateRequest(ctx echo.Context, request interface{}) (result interface{}, err error)

	ExtractValidationErrors(errs validator.ValidationErrors) (errMsg string)
}

var lock = &sync.Mutex{}
var reqHelper *RequestHelper

//GetInstance is used to get Request Helper
func GetInstance() IRequestHelper {
	if reqHelper == nil {
		lock.Lock()
		defer lock.Unlock()
		if reqHelper == nil {
			en := en.New()
			uni := ut.New(en)

			reqHelper = new(RequestHelper)

			reqHelper.translator, _ = uni.GetTranslator("en")
			reqHelper.validator = validator.New()

			err := enTranslations.RegisterDefaultTranslations(reqHelper.validator, reqHelper.translator)
			exception.PanicIfNeeded(err)
		}
	}

	return reqHelper
}

//RequestHelper struct
//Request Helper is used to validate input body
//Access https://pkg.go.dev/github.com/go-playground/validator/v10@v10.3.0?tab=doc
//for complete documentations
type RequestHelper struct {
	validator  *validator.Validate
	translator ut.Translator
}

//ValidateRequest is used to validate request body
func (rh *RequestHelper) ValidateRequest(ctx echo.Context, request interface{}) (result interface{}, err error) {
	if err = ctx.Bind(request); err != nil {
		return nil, err
	}

	if err = rh.validator.Struct(request); err != nil {
		return nil, errors.New(rh.ExtractValidationErrors(err.(validator.ValidationErrors)))
	}

	return request, nil
}

//ExtractValidationErrors is used to extract validation errors
func (rh *RequestHelper) ExtractValidationErrors(errs validator.ValidationErrors) (errMsg string) {

	errTranslation := errs.Translate(rh.translator)

	for _, err := range errs {
		jsonField := str.Underscore(err.Field())
		errMsg += strings.Replace(errTranslation[err.Namespace()], err.StructField(), jsonField, -1) + "; "
	}

	return errMsg
}
