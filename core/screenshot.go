package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/nahid/gohttp"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/utils"
	"io/ioutil"
)

func TakeScreenShotAndSave(url string, directory string) error {
	chromeUrl, err := getDebugURL()
	if err != nil {
		return err
	}

	orgCtx, cancelCtx := chromedp.NewRemoteAllocator(context.Background(), chromeUrl)
	ctx, _ := chromedp.NewContext(orgCtx)
	defer cancelCtx()

	url = utils.FormatUrlWithProtocol(url)

	var result []byte

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&result),
	}); err != nil {
		return err
	}

	url = utils.FormatUrlWithoutProtocol(url)

	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.png", directory, url), result, 0644); err != nil {
		return err
	}
	return nil
}

func TakeScreenShot(url string) ([]byte, error) {
	chromeUrl, err := getDebugURL()
	if err != nil {
		return nil, err
	}

	orgCtx, cancelCtx := chromedp.NewRemoteAllocator(context.Background(), chromeUrl)
	ctx, _ := chromedp.NewContext(orgCtx)
	defer cancelCtx()

	url = utils.FormatUrlWithProtocol(url)

	var result []byte

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&result),
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func getDebugURL() (string, error) {
	resp, err := gohttp.NewRequest().
		Get(fmt.Sprintf("%s/json/version", config.App().ChromeHeadlessUrl))
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.GetBody())
	if err != nil {
		return "", err
	}

	var result map[string]interface{}

	if err := json.Unmarshal(b, &result); err != nil {
		return "", err
	}
	return result["webSocketDebuggerUrl"].(string), nil
}
