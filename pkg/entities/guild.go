package entities

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) CreateGuild(guild request.CreateGuildRequest) error {
	err := e.repo.DiscordGuilds.CreateIfNotExists(model.DiscordGuild{
		ID:   guild.ID,
		Name: guild.Name,
		BotScopes: model.JSONArrayString{
			"*",
		},
	})
	if err != nil {
		return err
	}

	// notifiy new guild to discord
	err = e.svc.Discord.NotifyNewGuild(guild.ID)
	if err != nil {
		e.log.Errorf(err, "failed to notify new guild %s to discord", guild.ID)
	}

	return nil
}

func (e *Entity) GetGuilds() (*response.GetGuildsResponse, error) {
	guilds, err := e.repo.DiscordGuilds.Gets()
	if err != nil {
		return nil, err
	}

	var res response.GetGuildsResponse
	res.Data = make([]*response.GetGuildResponse, 0)
	for _, g := range guilds {
		res.Data = append(res.Data, &response.GetGuildResponse{
			ID:           g.ID,
			Name:         g.Name,
			BotScopes:    g.BotScopes,
			Alias:        g.Alias,
			LogChannelID: g.GuildConfigInviteTracker.ChannelID,
		})
	}

	return &res, nil
}

func (e *Entity) GetGuild(guildID string) (*response.GetGuildResponse, error) {
	guild, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &response.GetGuildResponse{
		ID:           guild.ID,
		Name:         guild.Name,
		BotScopes:    guild.BotScopes,
		Alias:        guild.Alias,
		LogChannelID: guild.GuildConfigInviteTracker.ChannelID,
		GlobalXP:     guild.GlobalXP,
	}, nil
}

func listDiscordGuilds(s *discordgo.Session) ([]*discordgo.UserGuild, error) {

	var (
		guilds  []*discordgo.UserGuild
		afterID string
	)

	for {
		tmp, err := s.UserGuilds(100, "", afterID)
		if err != nil {
			return nil, err
		}

		afterID = tmp[len(tmp)-1].ID
		guilds = append(guilds, tmp...)

		if len(tmp) < 100 {
			break
		}
	}

	return guilds, nil
}

type DiscordGuildResponse struct {
	discordgo.UserGuild
	BotAddable bool `json:"bot_addable"`
	BotArrived bool `json:"bot_arrived"`
}

func (e *Entity) ListMyDiscordGuilds(accessToken string) ([]*DiscordGuildResponse, error) {
	s, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to open discord session: %v", err.Error())
	}

	userGuilds, err := listDiscordGuilds(s)
	if err != nil {
		return nil, fmt.Errorf("failed to list user's discord guilds: %v", err.Error())
	}

	mochiGuilds, err := listDiscordGuilds(e.discord)
	if err != nil {
		return nil, fmt.Errorf("failed to list mochi's discord guilds: %v", err.Error())
	}

	mochiArrived := make(map[string]bool)

	for _, g := range mochiGuilds {
		mochiArrived[g.ID] = true
	}

	res := make([]*DiscordGuildResponse, 0)
	for _, g := range userGuilds {
		// Check for guilds that user has ADMINISTRATOR or MANAGE_GUILD permission
		if (g.Permissions&0x8) == 0x8 || (g.Permissions&0x20) == 0x20 {
			res = append(res, &DiscordGuildResponse{*g, true, mochiArrived[g.ID]})
		}
	}

	return res, nil
}

func (e *Entity) ToggleGuildGlobalXP(guildID string, globalXP bool) error {
	if err := e.repo.DiscordGuilds.ToggleGlobalXP(guildID, globalXP); err != nil {
		e.log.Errorf(err, "failed to toggle global XP for guild %s", guildID)
		return err
	}
	return nil
}
