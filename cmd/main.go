package main

import (
	"TwitterPrometheusExporter/exporter"
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
	consumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken    = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret   = os.Getenv("TWITTER_ACCESS_SECRET")

	addr = flag.String("port", "8081", "The port to listen on for HTTP requests.")
	user = flag.String("user", "", "Twitter account name")
	path = flag.String("path", "/metrics", "Metrics path")
)

var profile *exporter.TwitterProfile

func main() {
	flag.Parse()

	fmt.Printf("Looking up Twitter user '%s'\n", *user)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	if len(consumerKey) == 0 {
		panic("TWITTER_CONSUMER_KEY not set")
	}
	if len(consumerSecret) == 0 {
		panic("TWITTER_CONSUMER_SECRET not set")
	}
	if len(accessToken) == 0 {
		panic("TWITTER_ACCESS_TOKEN not set")
	}
	if len(accessSecret) == 0 {
		panic("TWITTER_ACCESS_SECRET not set")
	}

	token := oauth1.NewToken(accessToken, accessSecret)

	profile = exporter.NewTwitterProfile(*user, config.Client(oauth1.NoContext, token))

	collector := exporter.NewCollector(profile)
	prometheus.MustRegister(collector)

	http.Handle(*path, promhttp.Handler())
	fmt.Printf("Metrics available on http://localhost:%s%s\n", *addr, *path)

	port := fmt.Sprintf(":%v", *addr)
	log.Fatal(http.ListenAndServe(port, nil))
}
