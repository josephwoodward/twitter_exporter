package exporter

import (
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
)

type TwitterProfile struct {
	ScreenName string
	client     twitter.Client
}

func NewProfile(twitterHandle string, client *http.Client) *TwitterProfile {
	return &TwitterProfile{
		client:     *twitter.NewClient(client),
		ScreenName: twitterHandle,
	}
}

func (p *TwitterProfile) totalLikes() (int, error) {
	yes := true
	no := false

	tweets, _, err := p.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      p.ScreenName,
		Count:           100,
		ExcludeReplies:  &yes,
		IncludeRetweets: &no,
	})

	if err != nil {
		return -1, nil
	}

	var totalLikes int
	for _, v := range tweets {
		totalLikes += v.FavoriteCount
	}

	return totalLikes, nil
}

func (p *TwitterProfile) FetchUser() (*twitter.User, error) {
	user, _, err := p.client.Users.Show(&twitter.UserShowParams{
		ScreenName: p.ScreenName,
	})

	return user, err
}
