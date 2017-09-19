package main

import (
	"fmt"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

func getPrefix(zabbixPrefix, key string) string {
	return fmt.Sprintf("%s.%s", zabbixPrefix, key)
}

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
				getPrefix(zabbixPrefix, statName),
				stat,
			),
		)
	}

	return zabbixMetrics
}

func createOpcacheMetrics(
	opcacheStats *Opcache,
	hostname,
	zabbixPrefix string,
	zabbixMetrics []*zsend.Metric,
) []*zsend.Metric {
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.opcache_enabled"),
			strconv.FormatBool(opcacheStats.OpcacheEnabled),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.cache_full"),
			strconv.FormatBool(opcacheStats.CacheFull),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.restart_pending"),
			strconv.FormatBool(opcacheStats.RestartPending),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.restart_in_progress"),
			strconv.FormatBool(opcacheStats.RestartInProgress),
		),
	)

	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.memory_usage.used_memory"),
			strconv.Itoa(int(opcacheStats.MemoryUsage.UsedMemory)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.memory_usage.free_memory"),
			strconv.Itoa(int(opcacheStats.MemoryUsage.FreeMemory)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(zabbixPrefix, "opcache.memory_usage.wasted_memory"),
			strconv.Itoa(int(opcacheStats.MemoryUsage.WastedMemory)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.memory_usage.current_wasted_percentage",
			),
			strconv.FormatFloat(
				opcacheStats.MemoryUsage.CurrentWastedPercentage,
				'f',
				-1,
				64,
			),
		),
	)

	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.interned_strings_usage.buffer_size",
			),
			strconv.Itoa(int(opcacheStats.InternedStringsUsage.BufferSize)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.interned_strings_usage.used_memory",
			),
			strconv.Itoa(int(opcacheStats.InternedStringsUsage.UsedMemory)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.interned_strings_usage.free_memory",
			),
			strconv.Itoa(int(opcacheStats.InternedStringsUsage.FreeMemory)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.interned_strings_usage.number_of_strings",
			),
			strconv.Itoa(int(opcacheStats.InternedStringsUsage.NumberOfStrigns)),
		),
	)

	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.num_cached_scripts",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.NumCachedSripts)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.num_cached_keys",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.NumCachedKeys)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.max_cached_keys",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.MaxCachedKeys)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.hits",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.Hits)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.start_time",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.StartTime)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.last_restart_time",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.LastRestartTime)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.oom_restarts",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.OomRestarts)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.hash_restarts",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.HashRestarts)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.manual_restarts",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.ManualRestarts)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.misses",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.Misses)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.blacklist_misses",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.BlacklistMisses)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.blacklist_missis_ratio",
			),
			strconv.Itoa(int(opcacheStats.OpcacheStatistics.BlacklistMissRatio)),
		),
	)
	zabbixMetrics = append(
		zabbixMetrics,
		zsend.NewMetric(
			hostname,
			getPrefix(
				zabbixPrefix,
				"opcache.opcache_statistics.opcache_hit_rate",
			),
			strconv.FormatFloat(
				opcacheStats.OpcacheStatistics.OpcacheHitRate,
				'f',
				-1,
				64,
			),
		),
	)

	return zabbixMetrics
}
