package services

import (
	"github.com/minio/minio-go"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/config"
	"io"
)

func CreateMinioBucket() error {
	conn := app.Minio()
	cfg := config.Minio()

	ok, err := conn.BucketExists(cfg.Bucket)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	if err := conn.MakeBucket(cfg.Bucket, cfg.Location); err != nil {
		return err
	}
	return nil
}

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
