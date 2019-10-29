package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go"
	"io"
)

type response struct {
	Status int         `json:"-"`
	Title  string      `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func (r *response) ServerJSON(ctx echo.Context) error {
	if err := ctx.JSON(r.Status, r); err != nil {
		return err
	}
	return nil
}

func (r *response) ServerImageFromMinio(ctx echo.Context, object *minio.Object) error {
	s, _ := object.Stat()
	fileName := fmt.Sprintf("%s.png", s.Key)
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	ctx.Response().Header().Set("Content-Type", s.ContentType)

	if _, err := io.Copy(ctx.Response().Writer, object); err != nil {
		return err
	}
	return nil
}
