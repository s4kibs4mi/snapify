package dao

import (
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/apimodels"
	"github.com/s4kibs4mi/snapify/ent"
)

type IScreenshotDao interface {
	Tx(tx *ent.Tx) (IScreenshotDao, *ent.Tx, error)
	Create(req *apimodels.ReqScreenshotCreate) (*ent.Screenshot, error)
	Get(ID uuid.UUID) (*ent.Screenshot, error)
	List(page, limit int) ([]*ent.Screenshot, error)
	Delete(ID uuid.UUID) error
}
