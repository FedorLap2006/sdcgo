package sdcgo

// UserRates contains all guilds which specific user gave a rate to. The key is guild id.
type UserRates map[string]GuildRate

// UserWarns contains information about user warns in blacklist system.
type UserWarns struct {
	UserID string `json:"id"`
	Type string `json:"type"`
	WarnsCount uint `json:"warns"`
}

