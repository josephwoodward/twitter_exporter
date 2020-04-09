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

	tweetsCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "twitter_user",
		Name:        "tweet_likes_total",
		ConstLabels: map[string]string{ "profile" : *user },
	}, []string{"id"})
)

var profile *TwitterProfile

func init() {
	prometheus.Register(totalFollowers)
	prometheus.Register(totalFollowing)
	prometheus.Register(totalTweets)

	prometheus.Register(tweetsCount)

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

func handler(w http.ResponseWriter, r *http.Request) {


	//for _, v := range profile.Tweets() {
	//	tweetsCount.WithLabelValues(v.IDStr).Add(float64(v.FavoriteCount))
	//}

}

func main() {

	flag.Parse()

	fmt.Printf("Looking up Twitter user '%s'\n", *user)
	fmt.Printf("Metrics available on http://localhost%s/metrics\n", *addr)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(token, tokenSecret)

	profile = NewProfile(*user, config.Client(oauth1.NoContext, token))


	flag.Parse()
	http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
		promhttp.Handler()
	})
	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8081", nil))

	//http.Handle("/metrics", promhttp.Handler())
	s := &http.Server{Addr: ":8081"}
	go func() {
		log.Print(s.ListenAndServe())
	}()

}
