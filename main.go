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

	totalFollowers = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "followers_total",
		ConstLabels: map[string]string{ "profile":"joe_mighty" },
	}, func() float64 {
		return FollowerCount()
	})

	totalFollowing = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "following_total",
		ConstLabels: map[string]string{ "profile":"joe_mighty" },
	}, func() float64 {
		return FollowingCount()
	})

	totalTweets = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "tweets_total",
		ConstLabels: map[string]string{ "profile":"joe_mighty" },
	}, func() float64 {
		return TotalTweets()
	})
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

var consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
var consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
var token = os.Getenv("TWITTER_ACCESS_TOKEN")
var tokenSecret = os.Getenv("TWITTER_ACCESS_SECRET")

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

	profile = NewProfile("joe_mighty", config.Client(oauth1.NoContext, token))

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

}
