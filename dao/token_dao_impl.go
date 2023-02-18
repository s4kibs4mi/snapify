package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/ent"
	"github.com/s4kibs4mi/snapify/ent/token"
	"github.com/s4kibs4mi/snapify/log"
)

type tokenDao struct {
	CommonDao
	logger log.IAppLogger
}

func NewTokenDao(client *ent.Client, logger log.IAppLogger) ITokenDao {
	return &tokenDao{
		CommonDao: CommonDao{
			client:           client,
			screenshotClient: client.Screenshot,
			tokenClient:      client.Token,
		},
		logger: logger,
	}
}

func (d *tokenDao) Tx(tx *ent.Tx) (ITokenDao, *ent.Tx, error) {
	var err error
	if tx == nil {
		tx, err = d.client.Tx(context.Background())
		if err != nil {
			return nil, nil, err
		}
	}

	return &tokenDao{
		CommonDao: CommonDao{
			client:           d.client,
			screenshotClient: tx.Screenshot,
			tokenClient:      tx.Token,
		},
		logger: d.logger,
	}, tx, nil
}

func (d *tokenDao) Create(token string) (*ent.Token, error) {
	return d.tokenClient.Create().
		SetID(uuid.New()).
		SetToken(token).
		Save(context.Background())
}

func (d *tokenDao) Get(t string) (*ent.Token, error) {
	return d.tokenClient.Query().
		Where(
			token.TokenEQ(t),
		).
		Only(context.Background())
}
