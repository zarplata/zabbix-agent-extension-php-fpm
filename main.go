package main

import (
	"fmt"
	"os"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
)

var version = "[manual build]"

func main() {
	usage := `zabbix-agent-extension-php-fpm

Usage:
	zabbix-agent-extension-php-fpm [options] 

Options:
	-s --script <name>           Script name [default: /status]
	-f --filename <name>         File name [default: /status]
	-n --dial-network <type>     Type of dial networks [default: unix]
	-a --dial-address <address>  Dial address [default: /run/php-fpm/php-fpm.sock]
	-o --opcache <url>           Opcahe stats url
	-z --zabbix-host <zhost>     Hostname or IP address of zabbix server [default: 127.0.0.1]
	-p --zabbix-port <zport>     Port of zabbix server [default: 10051]
	--zabbix-prefix <prefix>     Add part of your prefix for key [default: None]
	-h --help                    Show this screen.
`
	args, _ := docopt.Parse(usage, nil, true, version, false)

	zabbixHost := args["--zabbix-host"].(string)
	zabbixPort, err := strconv.Atoi(args["--zabbix-port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	zabbixPrefix := args["--zabbix-prefix"].(string)
	if zabbixPrefix == "None" {
		zabbixPrefix = "php-fpm"
	} else {
		zabbixPrefix = fmt.Sprintf("%s.%s", zabbixPrefix, "php-fpm")
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fcgiScript := args["--script"].(string)
	fcgiFilename := args["--filename"].(string)

	dialNetwork := args["--dial-network"].(string)
	dialAddress := args["--dial-address"].(string)

	fcgiParams := make(map[string]string)
	fcgiParams["SCRIPT_FILENAME"] = fcgiFilename
	fcgiParams["SCRIPT_NAME"] = fcgiScript

	stats := getStats(dialNetwork, dialAddress, fcgiParams)

	var zabbixMetrics []*zsend.Metric
	zabbixMetrics = createMetrics(stats, hostname, zabbixPrefix, zabbixMetrics)

	if opcacheURL, ok := args["--opcache"].(string); ok {
		fcgiParams["SCRIPT_NAME"] = ""
		fcgiParams["QUERY_STRING"] = opcacheURL
		opcacheStats := getOpcacheStats(dialNetwork, dialAddress, fcgiParams)
		zabbixMetrics = createOpcacheMetrics(
			opcacheStats,
			hostname,
			zabbixPrefix,
			zabbixMetrics,
		)
	}

	packet := zsend.NewPacket(zabbixMetrics)
	sender := zsend.NewSender(
		zabbixHost,
		zabbixPort,
	)
	sender.Send(packet)

	fmt.Println("OK")
}
