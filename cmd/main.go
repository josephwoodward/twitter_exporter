package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"twitterexporter/exporter"

	"github.com/dghubble/oauth1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	consumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken    = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret   = os.Getenv("TWITTER_ACCESS_SECRET")
)

func main() {
	opts := exporter.GetDefaultExporterOptions()

	flag.StringVar(&opts.ListenAddress, "addr", exporter.DefaultListenAddress, "Network host to listen on.")
	flag.StringVar(&opts.ListenAddress, "a", exporter.DefaultListenAddress, "Network host to listen on.")
	flag.IntVar(&opts.ListenPort, "port", exporter.DefaultListenPort, "Port to listen on.")
	flag.IntVar(&opts.ListenPort, "p", exporter.DefaultListenPort, "Port to listen on.")
	flag.StringVar(&opts.MetricsPath, "path", exporter.DefaultScrapePath, "URL path from which to serve scrapes.")

	flag.StringVar(&opts.Username, "user", "", "Twitter account name")
	flag.StringVar(&opts.Username, "u", "", "Twitter account name")

	flag.Parse()

	if len(opts.Username) <= 0 {
		log.Fatal("'user' parameter is not set")
	}
	if len(consumerKey) <= 0 {
		log.Fatal("TWITTER_CONSUMER_KEY not set")
	}
	if len(consumerSecret) <= 0 {
		log.Fatal("TWITTER_CONSUMER_SECRET not set")
	}
	if len(accessToken) <= 0 {
		log.Fatal("TWITTER_ACCESS_TOKEN not set")
	}
	if len(accessSecret) <= 0 {
		log.Fatal("TWITTER_ACCESS_SECRET not set")
	}

	if err := run(opts); err != nil {
		log.Fatal(err)
	}
}

func run(opts *exporter.TwitterExporterOptions) error {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	profile := exporter.NewTwitterProfile(opts.Username, config.Client(oauth1.NoContext, token))

	collector := exporter.NewCollector(profile)
	prometheus.MustRegister(collector)

	// Start the exporter
	mux := http.NewServeMux()
	mux.Handle(opts.MetricsPath, promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, opts.MetricsPath, http.StatusMovedPermanently)
	})

	fmt.Printf("Metrics available on http://%s:%d%s\n", opts.ListenAddress, opts.ListenPort, opts.MetricsPath)

	addr := fmt.Sprintf("%s:%v", opts.ListenAddress, opts.ListenPort)
	if err := http.ListenAndServe(addr, mux); err != nil {
		return fmt.Errorf("cannot start exporter: %v", err)
	}

	return nil
}
