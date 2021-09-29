package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Spotify struct {
	client *client
}

func New() *Spotify {
	return &Spotify{
		client: NewClient(),
	}
}

type AccessTokenResponse struct {
	ClientId                         string `json:"clientId"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs int64  `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      bool   `json:"isAnonymous"`
}

type Friend struct {
	User struct {
		Uri  string `json:"uri"`
		Name string `json:"name"`
	} `json:"user"`
	Track struct {
		Uri      string `json:"uri"`
		Name     string `json:"name"`
		ImageUrl string `json:"imageUrl"`
		Artist   struct {
			Uri  string `json:"uri"`
			Name string `json:"name"`
		} `json:"artist"`
	} `json:"track"`
}

type FriendActivityResponse struct {
	Friends []Friend `json:"friends"`
}

func (s *Spotify) GetWebAccessToken(spdcCookie string) (*AccessTokenResponse, error) {
	req, err := s.client.CreateRequest("GET", "https://open.spotify.com/get_access_token?reason=transport&productType=web_player")

	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "sp_dc",
		Value: spdcCookie,
	})

	body, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	var r AccessTokenResponse
	return &r, json.Unmarshal(body, &r)
}

func (s *Spotify) GetFriendActivity(accessToken string) (*FriendActivityResponse, error) {
	req, err := s.client.CreateRequest("GET", "https://guc-spclient.spotify.com/presence-view/v1/buddylist")

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	body, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	var r FriendActivityResponse
	return &r, json.Unmarshal(body, &r)
}
