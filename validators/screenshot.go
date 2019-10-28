package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/snapify/errors"
)

type ReqCreateScreenshot struct {
	URLs []string `json:"urls"`
}

func ValidateCreateAddress(ctx echo.Context) (*ReqCreateScreenshot, error) {
	pld := ReqCreateScreenshot{}
	if err := ctx.Bind(&pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		return &pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}
