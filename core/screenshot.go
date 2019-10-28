package core

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/s4kibs4mi/snapify/utils"
	"io/ioutil"
	"math"
)

func TakeScreenShotAndSave(url string, directory string) error {
	ctx, _ := chromedp.NewContext(context.Background())

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
	ctx, _ := chromedp.NewContext(context.Background())

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
		return nil, err
	}
	return result, nil
}
