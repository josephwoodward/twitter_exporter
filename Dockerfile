FROM --platform=linux/amd64 golang:1.17-alpine
ADD twitter_exporter /
EXPOSE 8081
ENTRYPOINT ["/twitter_exporter"]