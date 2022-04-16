package token

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetBySymbol(symbol string) (model.Token, error)
	GetAllSupported() ([]model.Token, error)
	GetByAddress(address string, chainID int) (*model.Token, error)
}
