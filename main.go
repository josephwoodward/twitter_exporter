package main

import (
	"flag"
	"github.com/dghubble/oauth1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var (
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	token = os.Getenv("TWITTER_ACCESS_TOKEN")
	tokenSecret = os.Getenv("TWITTER_ACCESS_SECRET")

	addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")
	user = flag.String("twitter.user", "", "Twitter account name")
)

var (

	totalFollowers = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "followers_total",
		ConstLabels: map[string]string{ "profile" : *user },
	}, func() float64 {
		return FollowerCount()
	})

	totalFollowing = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "following_total",
		ConstLabels: map[string]string{ "profile" : *user },
	}, func() float64 {
		return FollowingCount()
	})

	totalTweets = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "tweets_total",
		ConstLabels: map[string]string{ "profile" : *user },
	}, func() float64 {
		return TotalTweets()
	})
)

var profile *TwitterProfile

func init() {
	prometheus.Register(totalFollowers)
	prometheus.Register(totalFollowing)
	prometheus.Register(totalTweets)
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func FollowerCount() float64 {
	return profile.Followers()
}

func FollowingCount() float64 {
	return profile.Following()
}

func TotalTweets() float64 {
	return profile.TotalTweetsCount()
}

func main() {

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(token, tokenSecret)

	profile = NewProfile(*user, config.Client(oauth1.NoContext, token))

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

}
