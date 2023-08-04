package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	metricsSubsystemHttpServer = "http_srv"
)

type IHttpServerMetrics interface {
	IncNbConnections()
	DecNbConnections()
	IncNotFound(path string, SupplierOldId)
	IncMethodNotAllowed(method, path string, SupplierOldId)
	IncBadRequest(path string, SupplierOldId)
	IncUnauthorized(path string, SupplierOldId)
	IncForbidden(path string, SupplierOldId)
	IncMethodNotAllowed(method, path string, SupplierOldId)
	IncRequestTimeout(path string, SupplierOldId)
    IncTooManyRequests(path string, SupplierOldId)
}

type NoHttpServerMetrics struct{}

func (m *NoHttpServerMetrics) IncNbConnections()                                      {}
func (m *NoHttpServerMetrics) DecNbConnections()                                      {}
func (m *NoHttpServerMetrics) IncNotFound(string, SupplierOldId)                      {}
func (m *NoHttpServerMetrics) IncMethodNotAllowed(method, path string, SupplierOldId) {}
func (m *NoHttpServerMetrics) IncBadRequest(path string, SupplierOldId)               {}
func (m *NoHttpServerMetrics) IncUnauthorized(path string, SupplierOldId)             {}
func (m *NoHttpServerMetrics) IncForbidden(path string, SupplierOldId)                {}
func (m *NoHttpServerMetrics) IncRequestTimeout(path string, SupplierOldId)           {}
func (m *NoHttpServerMetrics) IncTooManyRequests(path string, SupplierOldId)          {}

func NewHttpServerMetrics() IHttpServerMetrics {
	labels := map[string]string{
		metricsLabelMethod: "",
	}

	m := &httpServerMetrics{
		nbConnections: newGauge(metricsNamespace, metricsSubsystemHttpServer, "current_conns", nil),
		nbRequests:    newCounterVec(metricsNamespace, metricsSubsystemHttpServer, "nb_req", labels, []string{metricsLabelStatusCode}),
	}
	m.nbRequests400 = m.nbRequests.WithLabelValues("400")
	m.nbRequests401 = m.nbRequests.WithLabelValues("401")
	m.nbRequests403 = m.nbRequests.WithLabelValues("403")
	m.nbRequests404 = m.nbRequests.WithLabelValues("404")
	m.nbRequests405 = m.nbRequests.WithLabelValues("405")
	m.nbRequests408 = m.nbRequests.WithLabelValues("408")
	m.nbRequests429 = m.nbRequests.WithLabelValues("429")
	return m
}

type httpServerMetrics struct {
	nbConnections prometheus.Gauge
	nbRequests    *prometheus.CounterVec
	nbRequests400 prometheus.Counter
	nbRequests401 prometheus.Counter
	nbRequests403 prometheus.Counter
	nbRequests404 prometheus.Counter
	nbRequests405 prometheus.Counter
	nbRequests408 prometheus.Counter
	nbRequests429 prometheus.Counter
}

func (m *httpServerMetrics) IncNbConnections() {
	m.nbConnections.Inc()
}
func (m *httpServerMetrics) DecNbConnections() {
	m.nbConnections.Dec()
}

func (m *httpServerMetrics) IncNotFound(string) {
	m.nbRequests404.Inc()
}
func (m *httpServerMetrics) IncMethodNotAllowed(_, _ string) {
	m.nbRequests405.Inc()
}

func (m *httpServerMetrics) IncBadRequest(_, string) {
	m.nbRequests400.Inc()
}
func (m *httpServerMetrics) IncUnauthorized(_, string) {
	m.nbRequests401.Inc()
}

func (m *httpServerMetrics) IncForbidden(_, string) {
	m.nbRequests403.Inc()
}
func (m *httpServerMetrics) IncRequestTimeout(_, string) {
	m.nbRequests408.Inc()
}

func (m *httpServerMetrics) IncTooManyRequests(_, string) {
	m.nbRequests429.Inc()
}