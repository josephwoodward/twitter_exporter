package exporter

import (
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
)

type TwitterProfile struct {
	ScreenName string
	client     twitter.Client
}

func NewTwitterProfile(twitterHandle string, client *http.Client) *TwitterProfile {
	return &TwitterProfile{
		client:     *twitter.NewClient(client),
		ScreenName: twitterHandle,
	}
}

////TODO: Switch to polling to reduce unnecessary calls to Twitter
//func (p *TwitterProfile) Poll() {
//	ticker := time.NewTicker(5 * time.Second)
//	done := make(chan bool)
//
//	go func() {
//		for {
//			select {
//			case <-done:
//				return
//			case t := <-ticker.C:
//				//if user, err := p.fetchUser(); err == nil {
//				//}
//				//if timeline, err := p.fetchTimeline(); err == nil {
//				//}
//			}
//		}
//	}()
//}

func (p *TwitterProfile) fetchTimeline() (*timelineData, error) {
	//yes := true
	no := false

	tweets, _, err := p.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      p.ScreenName,
		Count:           50,
		ExcludeReplies:  &no,
		IncludeRetweets: &no,
	})

	if err != nil {
		return nil, err
	}

	t := &timelineData{}
	for _, v := range tweets {
		if !v.Retweeted {
			t.totalLikes += v.FavoriteCount
			t.totalRetweets += v.RetweetCount
			t.totalReplies += v.ReplyCount
		}
	}

	return t, nil
}

func (p *TwitterProfile) fetchUser() (*twitter.User, error) {
	user, _, err := p.client.Users.Show(&twitter.UserShowParams{
		ScreenName: p.ScreenName,
	})

	return user, err
}

type timelineData struct {
	totalLikes    int
	totalRetweets int
	totalReplies  int
}
