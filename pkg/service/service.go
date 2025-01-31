package service

import (
	"fmt"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/discord"
)

// import "github.com/defipod/api/pkg/service/binance"

type Service struct {
	CoinGecko coingecko.Service
	Discord   discord.Service
}

func NewService(
	cfg config.Config,
) (*Service, error) {

	discordSvc, err := discord.NewService(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	return &Service{
		CoinGecko: coingecko.NewService(),
		Discord:   discordSvc,
	}, nil
}
