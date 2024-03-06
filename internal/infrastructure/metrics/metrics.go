package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Requests = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "coupon_api_requests",
	Help: "Count of all requests",
})

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	// Prometheus Register
	prometheus.MustRegister(Requests)
}

func PrometheusMiddlewareGlobalCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/metrics" {
			Requests.Inc()
		}

		c.Next()
	}
}
