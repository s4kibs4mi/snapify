package core

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"math"
	"strings"
)

func TakeScreenShot(url string, directory string) error {
	ctx, _ := chromedp.NewContext(context.Background())

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

	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "http://", "", -1)
	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.png", directory, url), result, 0644); err != nil {
		return err
	}
	return nil
}
