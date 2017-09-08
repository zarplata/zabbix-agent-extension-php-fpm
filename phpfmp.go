package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tomasen/fcgi_client"
)

func getStats(
	dialNetwork, dialAddress string,
	fcgiParams map[string]string,
) map[string]string {
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

	statsBody, err := ioutil.ReadAll(statsResponse.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	stats := parseStats(string(statsBody))

	return stats
}

func parseStats(statsBody string) map[string]string {
	stats := make(map[string]string)

	statsByLine := strings.Split(statsBody, "\n")

	for _, statistic := range statsByLine {
		if statistic == "" {
			continue
		}

		splitStatistc := strings.Split(statistic, ":")

		statName := strings.Replace(splitStatistc[0], " ", "_", -1)
		statValue := strings.TrimSpace(strings.Join(splitStatistc[1:], ":"))

		stats[statName] = statValue
	}

	return stats
}
