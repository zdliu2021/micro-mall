package model

import "time"

type OAuthGitteResponse struct {
	ID                int         `json:"id"`
	Login             string      `json:"login"`
	Name              string      `json:"name"`
	AvatarURL         string      `json:"avatar_url"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	Remark            string      `json:"remark"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	Blog              interface{} `json:"blog"`
	Weibo             interface{} `json:"weibo"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	Stared            int         `json:"stared"`
	Watched           int         `json:"watched"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	Email             interface{} `json:"email"`
}
