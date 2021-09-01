package main

import (
	"dnck.github.io/tools/dns-tool/server"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var settings dnsserver.Settings

func init() {
	flag.StringVar(&settings.Address,
		"address",
		"0.0.0.0",
		"ipv4 address to listen on")
	flag.IntVar(&settings.Port,
		"port",
		5353,
		"port to bind to")
	flag.StringVar(&settings.DnsResolverPin,
		"pin",
		"",
		"the base64 encoded sha256 hash of the trusted dns resolver's tls cert (SPKI)")
	flag.Parse()
}

func main() {
	err := settings.MakeDnsProxyServer(); if err != nil {
		fmt.Println(err.Error())
	}
	// Wait for SIGINT or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

}
