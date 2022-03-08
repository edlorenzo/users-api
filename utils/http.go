package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/edlorenzo/users-api/logs"
)

var logger = logs.Logs

type Conf struct {
	UsersURI string `mapstructure:"usersURI"`
}

type UserResponseData struct {
	Data []*Data `json:"data"`
}

type Data struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func UserListHttpHandler(cfg *Conf) (error, bool, []byte) {
	resp, err := http.Get(cfg.UsersURI)
	if err != nil {
		logger.Errorf("[error in get url] %s", err.Error())
		return err, false, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("[error response.Body] %s", err.Error())
		return err, false, nil
	}

	var users UserResponseData
	err = json.Unmarshal(body, &users)
	if err != nil {
		logger.Errorf("[error json.Unmarshal] %s", err.Error())
		return err, false, nil
	}

	if resp.StatusCode == 200 {
		return nil, true, body
	} else {
		return err, false, body
	}
}

func HttpConnectionStatus(cfg *Conf) (bool, error) {
	netTransport := &http.Transport{
		MaxConnsPerHost:       10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	res, err := httpClient.Get(cfg.UsersURI)
	if err != nil {
		logger.Errorf("[error http transport] %s", err.Error())
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return false, err
	}

	return true, nil
}
