package validators

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/snapify/errors"
)

type ReqCreateScreenshot struct {
	URLs []string `json:"urls" valid:"required;range(1|10000)"`
}

func ValidateCreateScreenshot(ctx echo.Context) (*ReqCreateScreenshot, error) {
	pld := ReqCreateScreenshot{}
	if err := ctx.Bind(&pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		ve := errors.ValidationError{}

		for _, s := range pld.URLs {
			if !govalidator.IsURL(s) {
				ve.Add(fmt.Sprintf("url: %s", s), "is invalid")
			}
		}

		if len(ve) > 0 {
			return nil, &ve
		}
		return &pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}
