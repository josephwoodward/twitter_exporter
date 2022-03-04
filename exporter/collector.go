package exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DefaultListenPort    = 8081
	DefaultMaxTweets     = 50
	DefaultListenAddress = "0.0.0.0"
	DefaultScrapePath    = "/metrics"
)

type TwitterExporterOptions struct {
	ListenAddress string
	ListenPort    int
	MetricsPath   string
	Username      string
	MaxTweets     int
}

func GetDefaultExporterOptions() *TwitterExporterOptions {
	opts := &TwitterExporterOptions{
		ListenAddress: DefaultListenAddress,
		ListenPort:    DefaultListenPort,
		MetricsPath:   DefaultScrapePath,
	}

	return opts
}

var _ prometheus.Collector = &Collector{}

type Collector struct {
	twitter        *TwitterProfile
	maxTweets      int
	totalTweets    *prometheus.Desc
	totalFollowers *prometheus.Desc
	totalFollowing *prometheus.Desc
	totalLikes     *prometheus.Desc
	totalRetweets  *prometheus.Desc
	totalReplies   *prometheus.Desc
}

func NewCollector(profile *TwitterProfile, maxTweets int) *Collector {
	labels := map[string]string{"username": profile.ScreenName}
	return &Collector{
		twitter:   profile,
		maxTweets: maxTweets,
		totalTweets: prometheus.NewDesc(
			"twitter_profile_tweets_total",
			"Total number of tweets for user",
			nil,
			labels),
		totalFollowers: prometheus.NewDesc(
			"twitter_profile_followers_total",
			"Total number of followers for user",
			nil,
			labels),
		totalFollowing: prometheus.NewDesc(
			"twitter_profile_following_total",
			"Total following for user",
			nil,
			labels),
		totalLikes: prometheus.NewDesc(
			"twitter_timeline_likes_total",
			fmt.Sprintf("Total likes of profile tweets over last %v tweets", maxTweets),
			nil,
			labels),
		totalRetweets: prometheus.NewDesc(
			"twitter_timeline_retweets_total",
			fmt.Sprintf("Total retweets of profile tweets over last %v tweets", maxTweets),
			nil,
			labels),
		totalReplies: prometheus.NewDesc(
			"twitter_timeline_replies_total",
			fmt.Sprintf("Total replies of profile tweets over last %v tweets", maxTweets),
			nil,
			labels),
	}
}

func (c *Collector) Describe(d chan<- *prometheus.Desc) {
	d <- c.totalTweets
	d <- c.totalFollowers
	d <- c.totalFollowing
	d <- c.totalLikes
}

func (c *Collector) Collect(m chan<- prometheus.Metric) {
	if user, err := c.twitter.fetchUser(); err == nil {
		m <- prometheus.MustNewConstMetric(c.totalTweets, prometheus.GaugeValue, float64(user.StatusesCount))
		m <- prometheus.MustNewConstMetric(c.totalFollowers, prometheus.GaugeValue, float64(user.FollowersCount))
		m <- prometheus.MustNewConstMetric(c.totalFollowing, prometheus.GaugeValue, float64(user.FriendsCount))
	}

	if timeline, err := c.twitter.fetchTimeline(c.maxTweets); err == nil {
		m <- prometheus.MustNewConstMetric(c.totalLikes, prometheus.GaugeValue, float64(timeline.totalLikes))
		m <- prometheus.MustNewConstMetric(c.totalReplies, prometheus.GaugeValue, float64(timeline.totalReplies))
		m <- prometheus.MustNewConstMetric(c.totalRetweets, prometheus.GaugeValue, float64(timeline.totalRetweets))
	}
}
