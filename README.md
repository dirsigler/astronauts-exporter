# ISS Astronauts Exporter

This project contains a Prometheus Exporter written in Golang to learn the basic to create custom Exporter.
Goal is to fetch remote data and get the current number of Astronauts stationed on the ISS.

The total number of Astronauts is then exported as a custom Prometheus metric.

## Requirements

- Golang
- A working connection to - http://api.open-notify.org/astros.json

## Getting started

Clone the project into your local workspace.

Set a environment variable for the API endpoint we call `export NASA_ENDPOINT="http://api.open-notify.org/astros.json"`

Run the Golang project `go run main.go`

Open the Prometheus Metrics page of the project: `open http://localhost:9141/metrics`

Check the custom Prometheus Exporter Metric under the configured $Namespace+MetricName:
`my_custom_exporter_count_nasa_astronauts_total`

```sh
$ curl -sS http://localhost:9141/metrics | grep "my_custom_exporter_count_nasa_astronauts_total"
# HELP my_custom_exporter_count_nasa_astronauts_total Count of Astronauts currently on ISS.
# TYPE my_custom_exporter_count_nasa_astronauts_total counter
my_custom_exporter_count_nasa_astronauts_total 10
```

## Inspiration

- https://medium.com/teamzerolabs/15-steps-to-write-an-application-prometheus-exporter-in-go-9746b4520e26
- https://github.com/prometheus/haproxy_exporter/blob/d845c01ece3b922fba5813bf2773a2d43d6c2d4c/haproxy_exporter.go
- https://github.com/teamzerolabs/mirth_channel_exporter/blob/f4388395a0be997fe370537240a4da0fb746aec5/main.go#L151
- https://gist.github.com/fl64/a86b3d375deb947fa11099dd374660da
- https://reachmnadeem.wordpress.com/2021/01/06/writing-prometheus-exporter-using-go/
- https://www.youtube.com/watch?v=3wT0zSsQb58
- https://www.youtube.com/watch?v=2USCcDbbAZc
- https://percona.community/blog/2021/07/21/create-your-own-exporter-in-go/
- https://rsmitty.github.io/Prometheus-Exporters/
- https://dev.to/metonymicsmokey/custom-prometheus-metrics-with-go-520n
- https://stackoverflow.com/questions/73215835/golang-prometheus-exporter-pull-metrics-on-demand
