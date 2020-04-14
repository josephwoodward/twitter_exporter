package main

import "github.com/prometheus/client_golang/prometheus"

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
