package sdcgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (s *Session) request(method string, urlStr string, data url.Values) (response []byte, err error) {
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	if s.Token != "" {
		req.Header.Add("Authorization", s.Token)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = s.RateLimiter.Wait(context.Background())
	if err != nil {
		return
	}

	httpResp, err := s.Client.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(httpResp.Body)
	return
}

func unmarshal(b []byte, data interface{}) error {
	resterr, _ := NewRESTError(b)
	if resterr != nil {
		return resterr
	}

	return json.Unmarshal(b, data)
}


// Guild retrieves and returns all guild info.
func (s *Session) Guild(guildID string) (guild *Guild, err error) {
	response, err := s.request("GET", EndpointGuild(guildID), nil)
	if err != nil {
		return
	}
	err = unmarshal(response, &guild)
	return
}

// GuildPlaces returns all guild rates.
func (s *Session) GuildRates(guildID string) (rates GuildRates, err error) {
	response, err := s.request("GET", EndpointGuildRates(guildID), nil)
	if err != nil {
		return
	}

	err = unmarshal(response, &rates)
	return
}

// GuildRates returns place of the guild.
func (s *Session) GuildPlace(guildID string) (place uint, err error) {
	response, err := s.request("GET", EndpointGuildPlace(guildID), nil)
	if err != nil {
		return
	}

	var raw struct{
		Place uint `json:"place"`
	}
	err = unmarshal(response, &raw)
	if err != nil {
		return
	}

	return raw.Place, nil
}

// UserRates collects and returns all rates on different servers made by specific user.
func (s *Session) UserRates(userID string) (rates UserRates, err error) {
	response, err := s.request("GET", EndpointUserRates(userID), nil)
	if err != nil {
		return
	}

	err = unmarshal(response, &rates)
	return
}

// UserWarns returns all user blacklist warns.
func (s *Session) UserWarns(userID string) (warns *UserWarns, err error) {
	response, err := s.request("GET", EndpointBlacklistWarns(userID), nil)
	if err != nil {
		return
	}
	err = unmarshal(response, &warns)
	return
}

func (s *Session) PostStats(botID string, stats BotStats) (status bool, err error) {
	response, err := s.request("POST", EndpointBotStats(botID), url.Values {
		"servers": {strconv.Itoa(stats.Guilds)},
		"shards": {strconv.Itoa(stats.Shards)},
	})
	if err != nil {
		return
	}
	fmt.Println(string(response))
	var rawStatus struct {
		Status bool `json:"status"`
	}
	err = unmarshal(response, &rawStatus)
	if err != nil {
		return
	}
	return rawStatus.Status, nil
}