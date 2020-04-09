package main

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

func (profile *TwitterProfile) Handler() {

}

func (profile *TwitterProfile) Followers() float64 {
	user, _, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,

	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.FollowersCount)
}

func (profile *TwitterProfile) TotalTweetsCount() float64 {
	user, _, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.StatusesCount)
}

func (profile *TwitterProfile) Tweets() []twitter.Tweet {
	enabled := true
	disabled := false

	tweets, _, err := profile.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      profile.ScreenName,
		Count:           20,
		ExcludeReplies: &enabled,
		IncludeRetweets: &disabled,
	})

	if err != nil {
		fmt.Println(err)
	}

	return tweets
}

func (profile *TwitterProfile) Following() float64 {


	//user, _, err := profile.client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{})
	user, _, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return float64(user.FriendsCount)
}
