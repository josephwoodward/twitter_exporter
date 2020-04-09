FROM golang:1.14-alpine

ADD TwitterPrometheusExporter /

EXPOSE 8081

ENTRYPOINT ["/TwitterPrometheusExporter"]
