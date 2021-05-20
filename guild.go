package sdcgo

import (
	"encoding/json"
	"strings"
)

// GuildTags represents all guild tags.
type GuildTags []string

func (t *GuildTags) UnmarshalJSON(b []byte) error {
	v := ""
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	*t = strings.Split(v, ",")
	return nil
}

// GuildStatus is a status (badge) of a specific guild.
type GuildStatus int

// Guild statuses
const (
	SiteDevGuild GuildStatus = 1 << iota
	VerifiedGuild
	PartnerGuild
	FavoriteGuild
	BugHunterGuild
	EasterEggGuild
	BotDevGuild
	YoutubeGuild
	TwitchGuild
	SpamHuntGuild
)

// Guild contains all SDC specific information about a guild.
type Guild struct {
	Language   string      `json:"lang"`
	Name       string      `json:"name"`
	Desc       string      `json:"des"`
	Invite     string      `json:"invite"`
	Owner      string      `json:"owner"`
	Online     int         `json:"online"`
	Members    int         `json:"members"`
	Bot        int         `json:"bot"`
	BoostLevel int         `json:"boost"`
	Status     GuildStatus `json:"status"`
	UpCount    int         `json:"upCount"`
	Tags       GuildTags   `json:"tags"`
}

// GuildRate represents a rate given to a guild by an user.
type GuildRate int
// GuildRates contains all rates given to the guild. The key is user id.
type GuildRates map[string]GuildRate
