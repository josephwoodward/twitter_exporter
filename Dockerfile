FROM golang:1.17-alpine

ADD TwitterPrometheusExporter /

EXPOSE 8081

ENTRYPOINT ["/TwitterPrometheusExporter"]
