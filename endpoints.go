package sdcgo

// APIVersion is the latest supported version of the API.
var APIVersion = "2"

// Known SDC endpoints.
var (
	EndpointSDC = "https://api.server-discord.com/"
	EndpointAPI = EndpointSDC + "v" + APIVersion + "/"

	EndpointUsers     = EndpointAPI + "user/"
	EndpointUserRates = func(id string) string { return EndpointUsers + id + "/rated" }

	EndpointGuilds     = EndpointAPI + "guild/"
	EndpointGuild      = func(id string) string { return EndpointGuilds + id }
	EndpointGuildPlace = func(id string) string { return EndpointGuild(id) + "/place" }
	EndpointGuildRates = func(id string) string { return EndpointGuild(id) + "/rated" }

	EndpointBlacklistWarns = func(id string) string { return EndpointAPI + "warns/" + id }

	EndpointBots     = EndpointAPI + "bots/"
	EndpointBotStats = func(id string) string { return EndpointBots + id + "/stats" }
)
