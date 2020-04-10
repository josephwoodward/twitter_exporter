package main

import (
	"flag"
	"fmt"
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
	user = flag.String("user", "joe_mighty", "Twitter account name")
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

	tweetsCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "twitter_user",
		Name:        "tweet_likes_total",
		ConstLabels: map[string]string{ "profile" : *user },
	}, []string{"id"})
)


type Exporter struct {
	likesCount map[string]int
	tweets *prometheus.GaugeVec
}

func (e Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.tweets.Describe(ch)
}

func (e Exporter) Collect(ch chan<- prometheus.Metric) {
	for _, v := range profile.Tweets() {
		e.tweets.WithLabelValues(v.IDStr).Set(float64(v.FavoriteCount))
	}

	e.tweets.Collect(ch)
}

func NewExporter() *Exporter {
	e := Exporter{}
	e.tweets = tweetsCount
	e.likesCount = map[string]int{}

	// Count likes
	for _, v := range profile.Tweets() {
		e.likesCount[v.IDStr] = v.FavoriteCount
		e.tweets.WithLabelValues(v.IDStr).Set(float64(v.FavoriteCount))
	}

	return &e
}

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

	flag.Parse()

	fmt.Printf("Looking up Twitter user '%s'\n", *user)
	fmt.Printf("Metrics available on http://localhost%s/metrics\n", *addr)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(token, tokenSecret)

	profile = NewProfile(*user, config.Client(oauth1.NoContext, token))

	e := NewExporter()
	prometheus.MustRegister(e)

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))

}