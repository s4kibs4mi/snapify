package validators

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/snapify/errors"
	"net/http"
	"strings"
)

type ReqCreateScreenshot struct {
	URLs []string `json:"urls" valid:"required,length(1|10000)"`
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

func ValidateCreateScreenshotFromFile(ctx echo.Context) (*ReqCreateScreenshot, error) {
	if err := ctx.Request().ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	r := ctx.Request()
	r.Body = http.MaxBytesReader(ctx.Response(), r.Body, 32<<20) // 32 Mb

	f, h, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	body := make([]byte, h.Size)
	_, err = f.Read(body)
	if err != nil {
		return nil, err
	}

	pld := ReqCreateScreenshot{}

	ve := errors.ValidationError{}

	urls := strings.Split(string(body), ";")
	for _, u := range urls {
		u = strings.TrimSpace(u)
		if !govalidator.IsURL(u) {
			ve.Add(fmt.Sprintf("url: %s", u), "is invalid")
			continue
		}

		pld.URLs = append(pld.URLs, u)
	}

	if len(urls) > 10000 {
		ve.Add("urls", "must not be more than 10000")
	}

	if len(ve) > 0 {
		return nil, &ve
	}
	return &pld, nil
}
