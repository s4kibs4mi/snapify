package dao

import (
	"github.com/s4kibs4mi/snapify/apimodels"
	"github.com/s4kibs4mi/snapify/ent/enttest"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestScreenshotDao_ShouldCreateScreenshot(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	expectedURL := "https://example.com"

	dao := NewScreenshotDao(client, log.New())
	ss, err := dao.Create(&apimodels.ReqScreenshotCreate{
		URL: expectedURL,
	})
	require.NoError(t, err)
	require.NotNil(t, ss)
	require.NotEmpty(t, ss.ID.String())
	require.Equal(t, expectedURL, ss.URL)
}
