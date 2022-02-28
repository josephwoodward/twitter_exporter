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

	addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")
	user = flag.String("user", "_josephwoodward", "Twitter account name")
)

var profile *exporter.TwitterProfile

func main() {

	flag.Parse()

	fmt.Printf("Looking up Twitter user '%s'\n", *user)
	fmt.Printf("Metrics available on http://localhost%s/metrics\n", *addr)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	profile = exporter.NewTwitterProfile(*user, config.Client(oauth1.NoContext, token))

	collector := exporter.NewCollector(profile)
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
