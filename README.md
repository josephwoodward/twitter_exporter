# Twitter Prometheus Exporter

Why waste time writing meaningful instrumentation like a sucker when you can make your users do your
monitoring for you? They say that [the best alerting is done from the persepctive of the user](https://docs.google.com/document/d/199PqyG3UsyXlwieHaqbGiWVa8eMWi8zzAn0YfcApr8Q/),
so let's harness the fact that people are going to jump on Twitter to complain the moment something
goes wrong.

## Building

```bash
export GOOS="linux"
export GOARCH="amd64"

VERSION=$(git describe --always --dirty)
SHA1=$(git rev-parse --short --verify HEAD)
BUILD_DATE=$(date -u +%F-%T)

go build -ldflags "-extldflags -static -X main.VERSION=${VERSION} -X main.COMMIT_SHA1=${SHA1} -X main.BUILD_DATE=${BUILD_DATE}"
```

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

```bash
twitter_stream_exporter -twitter.track 'akeyword,anotherkeyword'
```

## Exported metrics

The exporter provides a set of counters that can be used to determine how frequently keywords are
being used.

| Metric | Notes |
| ------ | ----- |
| user_followers_total | The total number of followers the user has. |
| user_following_total | The total number of users the user is following |
| user_tweets_total | Total number of tweets |
|

A full sample of output can be found below.

```
user_followers_total{profile="joe_mighty"} 1479
user_following_total{profile="joe_mighty"} 639
user_tweets_total{profile="joe_mighty"} 11561
```
