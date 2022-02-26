package exporter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
)

func NewProfile(twitterHandle string, client *http.Client) *TwitterProfile {
	return &TwitterProfile{
		client:     *twitter.NewClient(client),
		ScreenName: twitterHandle,
	}
}

type TwitterProfile struct {
	ScreenName string
	client     twitter.Client
}

func (p *TwitterProfile) Followers() float64 {
	user, _, err := p.client.Users.Show(&twitter.UserShowParams{
		ScreenName: p.ScreenName,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.FollowersCount)
}

func (p *TwitterProfile) TotalTweetsCount() float64 {
	user, _, err := p.client.Users.Show(&twitter.UserShowParams{
		ScreenName: p.ScreenName,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.StatusesCount)
}

func (p *TwitterProfile) Tweets() []twitter.Tweet {
	enabled := true
	disabled := false

	tweets, _, err := p.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      p.ScreenName,
		Count:           20,
		ExcludeReplies:  &enabled,
		IncludeRetweets: &disabled,
	})

	if err != nil {
		fmt.Println(err)
	}

	return tweets
}

func (p *TwitterProfile) Following() float64 {

	//user, _, err := p.client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{})
	user, _, err := p.client.Users.Show(&twitter.UserShowParams{
		ScreenName: p.ScreenName,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.FriendsCount)
}
