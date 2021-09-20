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
		"",
		"address the proxy server listens on")
	flag.StringVar(&settings.Port,
		"port",
		"53",
		"port the proxy server binds to")
	flag.IntVar(&settings.TimeoutSeconds,
		"timeout",
		3,
		"global read/write deadline for proxy server and tls-server client")
	flag.StringVar(&settings.DnsResolverPin,
		"resolver-pin",
		"",
		"the base64 encoded sha256 hash of the trusted dns resolver's tls cert (SPKI)")
	flag.StringVar(&settings.DnsResolverAddress,
		"resolver-addr",
		"1.1.1.1:853",
		"the trusted dns resolver's ip address and port")
	flag.StringVar(&settings.DnsResolverFqdn,
		"resolver-fqdn",
		"cloudflare-dns.com",
		"the trusted dns resolver's common name")
	flag.BoolVar(&settings.Debug,
		"debug",
		false,
		"print the DNS query")
	flag.Parse()
}

func main() {
	err := settings.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Wait for SIGINT or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
