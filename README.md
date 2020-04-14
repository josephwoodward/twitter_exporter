# Twitter Prometheus Exporter

Have you ever wanted to plot Twitter followers increasing or decreasing over time and see if they correlate to an event such as a blog post or Tweet? I have, so I build this.

This Prometheus exporter will allow you to track and monitor user metrics such as follows/unfollows, tweets and more over time to give you an idea what content is increasing or decreasing your follower count.

## Usage

You'll need to [Generate a twitter access token pair](https://dev.twitter.com/oauth/overview/application-owner-access-tokens).
The exporter doesn't need to post to Twitter, so you should set the newly created application's 
permissions model to "Read only".

Once that's done, export the four keys as environment variables.

```bash
export TWITTER_ACCESS_TOKEN="..."
export TWITTER_ACCESS_SECRET="..."
export TWITTER_CONSUMER_KEY="..."
export TWITTER_CONSUMER_SECRET="..."
```

Then run the exporter, for instance:

```bashm
twitter_prometheus_exporter -user=BillGates
```

## Exported metrics

The exporter provides a set of counters that can be used to determine how frequently keywords are
being used.

| Metric | Notes |
| ------ | ----- |
|twitter_user_followers_total | The total number of followers the user has. |
|twitter_user_following_total | The total number of users the user is following |
|twitter_user_tweets_total | Total number of tweets |

A full sample of output can be found below.

```
twitter_user_followers_total{profile="BillGates"} 1479
twitter_user_following_total{profile="BillGates"} 639
twitter_user_tweets_total{profile="BillGates"} 11561
```
