package dao

import "github.com/s4kibs4mi/snapify/ent"

type CommonDao struct {
	client           *ent.Client
	screenshotClient *ent.ScreenshotClient
}
