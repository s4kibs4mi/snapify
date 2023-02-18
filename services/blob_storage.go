package services

import (
	"bytes"
	"github.com/minio/minio-go"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

type IBlobStorageService interface {
	Init() error
	Save(localFilePath string) (string, error)
	FetchUrl(storedPath string) (string, error)
}

type minioService struct {
	bucketName     string
	signedDuration time.Duration
	client         *minio.Client
	logger         log.IAppLogger
}

func NewMinioService(cfg *config.AppCfg, logger log.IAppLogger) (IBlobStorageService, error) {
	c, err := minio.New(cfg.BlobStorageCfg.BaseURL,
		cfg.BlobStorageCfg.Key, cfg.BlobStorageCfg.Secret, cfg.BlobStorageCfg.IsSecure)
	if err != nil {
		return nil, err
	}

	return &minioService{
		client:         c,
		bucketName:     cfg.BlobStorageCfg.Bucket,
		signedDuration: cfg.BlobStorageCfg.SignDuration,
		logger:         logger,
	}, nil
}

func (s *minioService) Init() error {
	exists, err := s.client.BucketExists(s.bucketName)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if err := s.client.MakeBucket(s.bucketName, ""); err != nil {
		return err
	}

	return nil
}

func (s *minioService) Save(ssPath string) (string, error) {
	data, err := ioutil.ReadFile(ssPath)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}

	fileNameParts := strings.Split(ssPath, "/")

	reader := bytes.NewReader(data)
	storedFileName := fileNameParts[len(fileNameParts)-1]

	_, err = s.client.PutObject(s.bucketName, storedFileName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentDisposition: "inline",
		ContentType:        "image/x-png",
	})
	if err != nil {
		s.logger.Error(err)
		return "", err
	}

	return storedFileName, nil
}

func (s *minioService) FetchUrl(storedPath string) (string, error) {
	signedUrl, err := s.client.PresignedGetObject(s.bucketName, storedPath, s.signedDuration, url.Values{})
	if err != nil {
		return "", err
	}

	return signedUrl.String(), nil
}
