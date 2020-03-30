# Twitter Prometheus Exporter

This Prometheus exporter will allow you to track and monitor key user metrics such as follows/unfollows, tweets and more.

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

Then run the exporter.

```bashm
twitter_prometheus_exporter -twitter.user ''
```

## Exported metrics

The exporter provides a set of counters that can be used to determine how frequently keywords are
being used.

| Metric | Notes |
| ------ | ----- |
|twitter_user_followers_total | The total number of followers the user has. |
|twitter_user_following_total | The total number of users the user is following |
|twitter_user_tweets_total | Total number of tweets |
|

A full sample of output can be found below.

```
twitter_user_followers_total{profile="joe_mighty"} 1479
twitter_user_following_total{profile="joe_mighty"} 639
twitter_user_tweets_total{profile="joe_mighty"} 11561
```
