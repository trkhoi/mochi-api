package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/activity"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	discordbottransaction "github.com/defipod/mochi/pkg/repo/discord_bot_transaction"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	"github.com/defipod/mochi/pkg/repo/token"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	"github.com/defipod/mochi/pkg/repo/users"
	whitelistcampaignusers "github.com/defipod/mochi/pkg/repo/whitelist_campaign_users"
	whitelistcampaigns "github.com/defipod/mochi/pkg/repo/whitelist_campaigns"
)

// NewRepo new pg repo implementation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{
		DiscordGuilds:                        discordguilds.NewPG(db),
		DiscordWalletVerification:            discordwalletverification.NewPG(db),
		InviteHistories:                      invitehistories.NewPG(db),
		Users:                                users.NewPG(db),
		UserWallet:                           userwallet.NewPG(db),
		GuildUsers:                           guildusers.NewPG(db),
		GuildCustomCommand:                   guildcustomcommand.NewPG(db),
		Token:                                token.NewPG(db),
		DiscordBotTransaction:                discordbottransaction.NewPG(db),
		DiscordUserGMStreak:                  discordusergmstreak.NewPG(db),
		GuildConfigGmGn:                      guildconfiggmgn.NewPG(db),
		GuildConfigInviteTracker:             guildconfiginvitetracker.NewPG(db),
		GuildConfigReactionRole:              guildconfigreactionrole.NewPG(db),
		GuildConfigDefaultRole:               guildconfigdefaultrole.NewPG(db),
		GuildConfigWalletVerificationMessage: guildconfigwalletverificationmessage.NewPG(db),
		DiscordGuildStats:                    discordguildstats.NewPG(db),
		DiscordGuildStatChannels:             discordguildstatchannels.NewPG(db),
		GuildConfigToken:                     guildconfigtoken.NewPG(db),
		WhitelistCampaigns:                   whitelistcampaigns.NewPG(db),
		WhitelistCampaignUsers:               whitelistcampaignusers.NewPG(db),
		NFTCollection:                        nftcollection.NewPG(db),
		Activity:                             activity.NewPG(db),
		GuildConfigActivity:                  guildconfigactivity.NewPG(db),
		ConfigXPLevel:                        configxplevel.NewPG(db),
		GuildUserActivityLog:                 guilduseractivitylog.NewPG(db),
		GuildUserXP:                          guilduserxp.NewPG(db),
		GuildConfigLevelRole:                 guildconfiglevelrole.NewPG(db),
	}
}
