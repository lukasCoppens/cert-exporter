package exporters

import (
	"github.com/joe-elliott/cert-exporter/src/metrics"
)

// CertExporter exports PEM file certs
type CertExporter struct {
}

// ExportMetrics exports the provided PEM file
func (c *CertExporter) ExportMetrics(file string) error {

	metric, err := secondsToExpiryFromCertAsFile(file)

	if err != nil {
		return err
	}

	metrics.CertExpirySeconds.WithLabelValues(file, metric.issuer, metric.cn).Set(metric.durationUntilExpiry)
	metrics.CertNotAfterTimestamp.WithLabelValues(file, metric.issuer, metric.cn).Set(metric.notAfter)
	return nil
}
