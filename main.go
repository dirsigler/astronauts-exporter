package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
)

const namespace = "my_custom_exporter"

type Astronauts struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	People  []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
}

type Exporter struct {
	URI string

	up               prometheus.Gauge
	numberAstronauts prometheus.Counter
}

func NewExporter(endpoint string) *Exporter {
	return &Exporter{
		URI: endpoint,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last scrape successful.",
		}),
		numberAstronauts: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "count_nasa_astronauts_total",
			Help:      "Count of Astronauts currently on ISS.",
		}),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.numberAstronauts.Desc()
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	req, err := http.NewRequest("GET", e.URI, nil)
	if err != nil {
		fmt.Printf("ERROR AT: %v", err.Error())
	}

	// Make request and show output.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR AT: %v", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("ERROR AT: %v", err.Error())
	}

	var astr Astronauts
	err = json.Unmarshal(body, &astr)
	if err != nil {
		fmt.Printf("ERROR AT: %v", err.Error())
	}
	ch <- prometheus.MustNewConstMetric(e.numberAstronauts.Desc(), prometheus.CounterValue, float64(astr.Number))
}

func main() {
	flag.Parse()

	endpoint := os.Getenv("NASA_ENDPOINT")

	exporter := NewExporter(endpoint)
	prometheus.MustRegister(exporter)
	log.Printf("Using connection endpoint: %s", endpoint)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>NASA ISS Exporter</title></head>
             <body>
             <h1>NASA ISS Exporter</h1>
             <p><a href='` + "/metrics" + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Fatal(http.ListenAndServe(":9141", nil))
}
