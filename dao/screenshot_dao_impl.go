package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/apimodels"
	"github.com/s4kibs4mi/snapify/ent"
	"github.com/s4kibs4mi/snapify/ent/screenshot"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/models"
	"time"
)

type screenshotDao struct {
	CommonDao
	logger log.IAppLogger
}

func NewScreenshotDao(client *ent.Client, logger log.IAppLogger) IScreenshotDao {
	return &screenshotDao{
		CommonDao: CommonDao{
			client:           client,
			screenshotClient: client.Screenshot,
			tokenClient:      client.Token,
		},
		logger: logger,
	}
}

func (d *screenshotDao) Tx(tx *ent.Tx) (IScreenshotDao, *ent.Tx, error) {
	var err error
	if tx == nil {
		tx, err = d.client.Tx(context.Background())
		if err != nil {
			return nil, nil, err
		}
	}

	return &screenshotDao{
		CommonDao: CommonDao{
			client:           d.client,
			screenshotClient: tx.Screenshot,
			tokenClient:      tx.Token,
		},
		logger: d.logger,
	}, tx, nil
}

func (d *screenshotDao) Create(req *apimodels.ReqScreenshotCreate) (*ent.Screenshot, error) {
	return d.screenshotClient.Create().
		SetID(uuid.New()).
		SetURL(req.URL).
		SetStatus(models.Queued).
		SetCreatedAt(time.Now()).
		Save(context.Background())
}

func (d *screenshotDao) Get(ID uuid.UUID) (*ent.Screenshot, error) {
	return d.screenshotClient.Get(context.Background(), ID)
}

func (d *screenshotDao) Delete(ID uuid.UUID) error {
	return d.screenshotClient.DeleteOneID(ID).Exec(context.Background())
}

func (d *screenshotDao) List(page, limit int) ([]*ent.Screenshot, error) {
	return d.screenshotClient.
		Query().
		Order(ent.Desc(screenshot.FieldCreatedAt)).
		Offset((page * limit) - limit).
		Limit(limit).
		All(context.Background())
}
