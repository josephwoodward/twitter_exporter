package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	user = flag.String("user", "", "Twitter account name")
)

func main() {
	opts := exporter.GetDefaultExporterOptions()

	flag.StringVar(&opts.ListenAddress, "addr", exporter.DefaultListenAddress, "Network host to listen on.")
	flag.StringVar(&opts.ListenAddress, "a", exporter.DefaultListenAddress, "Network host to listen on.")
	flag.IntVar(&opts.ListenPort, "port", exporter.DefaultListenPort, "Port to listen on.")
	flag.IntVar(&opts.ListenPort, "p", exporter.DefaultListenPort, "Port to listen on.")
	flag.StringVar(&opts.ScrapePath, "path", exporter.DefaultScrapePath, "URL path from which to serve scrapes.")

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

	if err := run(opts); err != nil {
		log.Fatal(err)
	}
}

func run(opts *exporter.TwitterExporterOptions) error {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	profile := exporter.NewTwitterProfile(*user, config.Client(oauth1.NoContext, token))

	collector := exporter.NewCollector(profile)
	prometheus.MustRegister(collector)

	// Start the exporter.
	http.Handle(opts.ScrapePath, promhttp.Handler())
	fmt.Printf("Metrics available on http://%s:%d%s\n", opts.ListenAddress, opts.ListenPort, opts.ScrapePath)

	p := fmt.Sprintf("%s:%v", opts.ListenAddress, opts.ListenPort)
	log.Fatal(http.ListenAndServe(p, nil))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	return nil
}
