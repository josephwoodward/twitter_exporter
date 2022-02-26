package exporter

import "github.com/prometheus/client_golang/prometheus"

type Collector struct {
	totalTweets    *prometheus.Desc
	totalLikes     *prometheus.Desc
	totalFollowers *prometheus.Desc
	totalFollowing *prometheus.Desc
}

func NewCollector(profile *TwitterProfile) *Collector {
	c := &Collector{
		totalTweets: prometheus.NewDesc(
			"tweets_total",
			"Total tweets for user",
			nil,
			nil),
		totalLikes: prometheus.NewDesc(
			"likes_total",
			"Total likes for user",
			nil,
			nil),
		totalFollowers: prometheus.NewDesc(
			"followers_total",
			"Total followers for user",
			nil,
			nil),
		totalFollowing: prometheus.NewDesc(
			"following_total",
			"Total following for user",
			nil,
			nil),
	}

	return c
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.totalTweets
	ch <- c.totalLikes
	ch <- c.totalFollowers
	ch <- c.totalFollowing
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {

	//TODO: Get twitter from collector data here

	ch <- prometheus.MustNewConstMetric(c.totalTweets, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.totalLikes, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.totalFollowers, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.totalFollowing, prometheus.GaugeValue, 1)
}
