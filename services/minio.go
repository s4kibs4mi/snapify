package services

import (
	"github.com/minio/minio-go"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/config"
	"io"
)

func UploadToMinio(fileName, contentType string, reader io.Reader, size int) error {
	conn := app.Minio()
	cfg := config.Minio()
	_, errP := conn.PutObject(cfg.Bucket, fileName, reader, int64(size), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if errP != nil {
		return errP
	}
	return nil
}

func ServeAsStreamFromMinio(fileName string) (*minio.Object, error) {
	conn := app.Minio()
	cfg := config.Minio()
	o, err := conn.GetObject(cfg.Bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return o, nil
}
