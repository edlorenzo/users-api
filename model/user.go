package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Title    string
	Content  string
	Author   string
	Creator  uint `gorm:"creator"`
	Modifier uint `gorm:"modifier"`
}

type NewUser []struct {
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

type Profile struct {
	Name        string `json:"name"`         // 1
	Login       string `json:"login"`        // 2
	Company     string `json:"company"`      // 3
	Followers   int    `json:"followers"`    // 4
	PublicRepos int    `json:"public_repos"` // 5
}

type ProfileList []struct {
	Name        string `json:"name"`         // 1
	Login       string `json:"login"`        // 2
	Company     string `json:"company"`      // 3
	Followers   int    `json:"followers"`    // 4
	PublicRepos int    `json:"public_repos"` // 5
}
