package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	addr              = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	uniformDomain     = flag.Float64("uniform.domain", 200, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 200, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 10, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "allen2",
			Help: "RPC latency distributions.",
		},
		[]string{"service"},
	)
	// The same as above, but now as a histogram, and only for the normal
	// distribution. The buckets are targeted to the parameters of the
	// normal distribution, with 20 buckets centered on the mean, each
	// half-sigma wide.
	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "allen",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
	})
)

// init方法自动调用, 建议每个package中只包含一个init
func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)
}

func main() {
	flag.Parse()

	start := time.Now()

	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	}

	// Periodically record some sample latencies for the three services.
	go func() {
		for {
			v := rand.Float64() * *uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * *normDomain) + *normMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			rpcDurationsHistogram.Observe(v)
			fmt.Println("rpcDurationsHistogram")
			time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)

		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64()
			rpcDurations.WithLabelValues("exponential").Observe(v)
			fmt.Println("rpcDurations")
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(*addr, nil)
}
