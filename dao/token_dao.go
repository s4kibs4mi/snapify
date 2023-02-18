package dao

import (
	"github.com/s4kibs4mi/snapify/ent"
)

type ITokenDao interface {
	Tx(tx *ent.Tx) (ITokenDao, *ent.Tx, error)
	Create(token string) (*ent.Token, error)
	Get(token string) (*ent.Token, error)
}
