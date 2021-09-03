package dnsserver

import (
	"fmt"
)

type Settings struct {
	Address            string
	Port               string
	TimeoutSeconds     int
	DnsResolverAddress string
	DnsResolverFqdn    string
	DnsResolverPin     string
}

func (settings *Settings) joinHostPort() string {
	return fmt.Sprintf("%s:%v", settings.Address, settings.Port)
}

func (settings *Settings) makeDnsClient() (*dnsClient, error) {
	dnsClient := dnsClient{settings.DnsResolverPin,
		settings.DnsResolverAddress,
		settings.DnsResolverFqdn,
		settings.TimeoutSeconds,
		nil,
	}
	err := dnsClient.establishTrust()
	if err != nil {
		return nil, err
	}
	return &dnsClient, nil
}

func (settings *Settings) makeDnsProxyServer() (*dnsProxyServer, error) {
	dnsClient, err := settings.makeDnsClient()
	if err != nil {
		return nil, err
	}
	return &dnsProxyServer{settings.Address, settings.Port,
		settings.TimeoutSeconds, dnsClient,
	}, nil
}

func (settings *Settings) Run() error {
	infof("dns proxy running @ tcp://localhost:5353")
	proxyServer, err := settings.makeDnsProxyServer()
	if err != nil {
		return err
	}
	return proxyServer.listenAndServe()
}
