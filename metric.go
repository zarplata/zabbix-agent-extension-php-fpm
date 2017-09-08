package main

import (
	"fmt"

	zsend "github.com/blacked/go-zabbix"
)

func createMetrics(
	stats map[string]string,
	hostname, zabbixPrefix string,
	zabbixMetrics []*zsend.Metric,
) []*zsend.Metric {
	for statName, stat := range stats {
		zabbixMetrics = append(
			zabbixMetrics,
			zsend.NewMetric(
				hostname,
				fmt.Sprintf("%s.%s", zabbixPrefix, statName),
				stat,
			),
		)
	}

	return zabbixMetrics
}
