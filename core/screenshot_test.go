package core

import (
	"testing"
)

func TestTakeScreenShot(t *testing.T) {
	out := "/Users/s4kibs4mi/go/src/github.com/s4kibs4mi/snapify/out"
	if err := TakeScreenShotAndSave("https://www.facebook.com", out); err != nil {
		t.Error(err)
		return
	}
	t.Log("Screen shot has been taken.")
}
