package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DefaultListenPort    = 8081
	DefaultListenAddress = "0.0.0.0"
	DefaultScrapePath    = "/metrics"
)

type TwitterExporterOptions struct {
	ListenAddress string
	ListenPort    int
	MetricsPath   string
	Username      string
}

func GetDefaultExporterOptions() *TwitterExporterOptions {
	opts := &TwitterExporterOptions{
		ListenAddress: DefaultListenAddress,
		ListenPort:    DefaultListenPort,
		MetricsPath:   DefaultScrapePath,
	}

	return opts
}

type Collector struct {
	twitter        *TwitterProfile
	totalTweets    *prometheus.Desc
	totalFollowers *prometheus.Desc
	totalFollowing *prometheus.Desc
	totalLikes     *prometheus.Desc
	totalRetweets  *prometheus.Desc
	totalReplies   *prometheus.Desc
}

func NewCollector(profile *TwitterProfile) prometheus.Collector {
	labels := map[string]string{"name": profile.ScreenName}
	return &Collector{
		twitter: profile,
		totalTweets: prometheus.NewDesc(
			"twitter_tweets_total",
			"Total number of tweets for user",
			nil,
			labels),
		totalFollowers: prometheus.NewDesc(
			"twitter_followers_total",
			"Total number of followers for user",
			nil,
			labels),
		totalFollowing: prometheus.NewDesc(
			"twitter_following_total",
			"Total following for user",
			nil,
			labels),
		totalLikes: prometheus.NewDesc(
			"twitter_likes_total",
			"Total likes for user in past 50 tweets",
			nil,
			labels),
		totalRetweets: prometheus.NewDesc(
			"twitter_retweets_total",
			"Total retweets for user",
			nil,
			labels),
		totalReplies: prometheus.NewDesc(
			"twitter_reply_total",
			"Total count of replies for user",
			nil,
			labels),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.totalTweets
	ch <- c.totalFollowers
	ch <- c.totalFollowing
	ch <- c.totalLikes
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	if user, err := c.twitter.fetchUser(); err == nil {
		ch <- prometheus.MustNewConstMetric(c.totalTweets, prometheus.GaugeValue, float64(user.StatusesCount))
		ch <- prometheus.MustNewConstMetric(c.totalFollowers, prometheus.GaugeValue, float64(user.FollowersCount))
		ch <- prometheus.MustNewConstMetric(c.totalFollowing, prometheus.GaugeValue, float64(user.FriendsCount))
	}

	if timeline, err := c.twitter.fetchTimeline(); err == nil {
		ch <- prometheus.MustNewConstMetric(c.totalLikes, prometheus.GaugeValue, float64(timeline.totalLikes))
		ch <- prometheus.MustNewConstMetric(c.totalReplies, prometheus.GaugeValue, float64(timeline.totalReplies))
		ch <- prometheus.MustNewConstMetric(c.totalRetweets, prometheus.GaugeValue, float64(timeline.totalRetweets))
	}
}
