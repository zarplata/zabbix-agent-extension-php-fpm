package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tomasen/fcgi_client"
)

type Opcache struct {
	OpcacheEnabled       bool                 `json:"opcache_enabled"`
	CacheFull            bool                 `json:"cache_full"`
	RestartPending       bool                 `json:"restart_pending"`
	RestartInProgress    bool                 `json:"restart_in_progress"`
	MemoryUsage          MemoryUsage          `json:"memory_usage"`
	InternedStringsUsage InternedStringsUsage `json:"interned_strings_usage"`
	OpcacheStatistics    OpcacheStatistics    `json:"opcache_statistics"`
}

type MemoryUsage struct {
	UsedMemory              int64   `json:"used_memory"`
	FreeMemory              int64   `json:"free_memory"`
	WastedMemory            int64   `json:"wasted_memory"`
	CurrentWastedPercentage float64 `json:"current_wasted_percentage"`
}

type InternedStringsUsage struct {
	BufferSize      int64 `json:"buffer_size"`
	UsedMemory      int64 `json:"used_memory"`
	FreeMemory      int64 `json:"free_memory"`
	NumberOfStrigns int64 `json:"number_of_strings"`
}

type OpcacheStatistics struct {
	NumCachedSripts    int64   `json:"num_cached_scripts"`
	NumCachedKeys      int64   `json:"num_cached_keys"`
	MaxCachedKeys      int64   `json:"max_cached_keys"`
	Hits               int64   `json:"hits"`
	StartTime          int64   `json:"start_time"`
	LastRestartTime    int64   `json:"last_restart_time"`
	OomRestarts        int64   `json:"oom_restarts"`
	HashRestarts       int64   `json:"hash_restarts"`
	ManualRestarts     int64   `json:"manual_restarts"`
	Misses             int64   `json:"misses"`
	BlacklistMisses    int64   `json:"blacklist_misses"`
	BlacklistMissRatio int64   `json:"blacklist_miss_ratio"`
	OpcacheHitRate     float64 `json:"opcache_hit_rate"`
}

func getOpcacheStats(
	dialNetwork,
	dialAddress string,
	fcgiParams map[string]string,
) *Opcache {
	var opcacheStats Opcache

	fcgi, err := fcgiclient.Dial(dialNetwork, dialAddress)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	statsResponse, err := fcgi.Get(fcgiParams)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.NewDecoder(statsResponse.Body).Decode(&opcacheStats)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return &opcacheStats
}
