package services

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"io/ioutil"
)

type IHeadlessBrowser interface {
	TakeScreenshot(url string) (string, error)
}

type headlessBrowserService struct {
	browserCon *rod.Browser
	logger     log.IAppLogger
}

func NewHeadlessBrowserService(cfg *config.AppCfg, logger log.IAppLogger) (IHeadlessBrowser, error) {
	lu, err := launcher.NewManaged(cfg.HeadlessBrowserUrl)
	if err != nil {
		return nil, err
	}

	browser := rod.New().Client(lu.MustClient())
	if err := browser.Connect(); err != nil {
		return nil, err
	}

	return &headlessBrowserService{
		browserCon: browser,
		logger:     logger,
	}, nil
}

func (s *headlessBrowserService) TakeScreenshot(url string) (string, error) {
	s.logger.Info("Browser taking screenshot: ", url)

	ssPath := fmt.Sprintf("/tmp/%s.png", uuid.New().String())
	page, err := s.browserCon.Page(proto.TargetCreateTarget{
		URL: url,
	})
	if err != nil {
		return "", err
	}

	if err := page.WaitLoad(); err != nil {
		return "", err
	}

	ssBytes, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
	})
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(ssPath, ssBytes, 0644); err != nil {
		return "", err
	}

	return ssPath, nil
}
