package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/nahid/gohttp"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/utils"
	"io/ioutil"
	"math"
)

func TakeScreenShotAndSave(url string, directory string) error {
	url, err := getDebugURL()
	if err != nil {
		return err
	}

	orgCtx, cancelCtx := chromedp.NewRemoteAllocator(context.Background(), url)
	ctx, _ := chromedp.NewContext(orgCtx)
	defer cancelCtx()

	url = utils.FormatUrlWithProtocol(url)

	var result []byte

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&result),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, view, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(view.Width)), int64(math.Ceil(view.Height))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			res, err := page.CaptureScreenshot().
				WithQuality(100).
				WithClip(&page.Viewport{
					X:      view.X,
					Y:      view.Y,
					Width:  view.Width,
					Height: view.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}

			result = res
			return nil
		}),
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
	url, err := getDebugURL()
	if err != nil {
		return nil, err
	}

	orgCtx, cancelCtx := chromedp.NewRemoteAllocator(context.Background(), url)
	ctx, _ := chromedp.NewContext(orgCtx)
	defer cancelCtx()

	url = utils.FormatUrlWithProtocol(url)

	var result []byte

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&result),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, viewHeight, viewWidth, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(viewWidth.Width)), int64(math.Ceil(viewHeight.ClientHeight))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			res, err := page.CaptureScreenshot().
				WithQuality(100).
				WithClip(&page.Viewport{
					X:      0,
					Y:      0,
					Width:  float64(width),
					Height: float64(height),
					Scale:  2,
				}).Do(ctx)
			if err != nil {
				return err
			}

			result = res
			return nil
		}),
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
