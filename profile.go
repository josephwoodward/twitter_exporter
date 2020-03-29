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

func (profile *TwitterProfile) Followers() float64 {
	user, resp, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,

	})

	if err != nil {
		fmt.Println(resp.Status)
		fmt.Print(user.Email)
	}

	return float64(user.FollowersCount)
}

func (profile *TwitterProfile) TotalTweetsCount() float64 {
	user, resp, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,
	})

	if err != nil {
		fmt.Println(resp.Status)
		fmt.Print(user.Email)
	}

	return float64(user.StatusesCount)
}

func (profile *TwitterProfile) Following() float64 {
	user, resp, err := profile.client.Users.Show(&twitter.UserShowParams{
		ScreenName: profile.ScreenName,
	})

	if err != nil {
		fmt.Println(resp.Status)
		fmt.Print(user.Email)
	}

	return float64(user.FriendsCount)
}
