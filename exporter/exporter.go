package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	totalFollowers = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "twitter_user",
		Name:      "followers_total",
	}, func() float64 {
		return FollowerCount()
	})

	totalFollowing = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "twitter_user",
		Name:      "following_total",
	}, func() float64 {
		return FollowingCount()
	})

	totalTweets = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "twitter_user",
		Name:      "tweets_total",
	}, func() float64 {
		return TotalTweets()
	})

	tweetsCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "twitter_user",
		Name:      "tweet_likes_total",
	}, []string{"id"})
)

func FollowerCount() float64 {
	return profile.Followers()
}

func FollowingCount() float64 {
	return profile.Following()
}

func TotalTweets() float64 {
	return profile.TotalTweetsCount()
}

type Collector struct {
	likesCount map[string]int
	tweets     *prometheus.GaugeVec
}

func NewCollector(profile *TwitterProfile) *Collector {
	c := Collector{
		likesCount: map[string]int{},
		tweets:     tweetsCount,
	}

	// Count likes
	for _, v := range profile.Tweets() {
		c.likesCount[v.IDStr] = v.FavoriteCount
		c.tweets.WithLabelValues(v.IDStr).Set(float64(v.FavoriteCount))
	}

	return &c
}

func (e *Collector) Describe(ch chan<- *prometheus.Desc) {
	e.tweets.Describe(ch)
}

func (e *Collector) Collect(ch chan<- prometheus.Metric) {
	for _, v := range e.profile.Tweets() {
		e.tweets.WithLabelValues(v.IDStr).Set(float64(v.FavoriteCount))
	}

	e.tweets.Collect(ch)
}
