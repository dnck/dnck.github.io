package dnsserver

import "fmt"

type Settings struct {
	Address        string
	Port           int
	DnsResolverPin string
}

func (settings *Settings) formatAddress() string {
	return fmt.Sprintf("%s:%v", settings.Address, settings.Port)
}
