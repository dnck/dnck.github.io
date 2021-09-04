package dnsserver

// Settings are explained in dns-tool.go
type Settings struct {
	Address            string
	Port               string
	TimeoutSeconds     int
	DnsResolverAddress string
	DnsResolverFqdn    string
	DnsResolverPin     string
}

// makeDnsClient creates and configures the tls server's client
func (settings *Settings) makeDnsClient() (*dnsClient, error) {
	dnsClient := dnsClient{settings.DnsResolverPin,
		settings.DnsResolverAddress,
		settings.DnsResolverFqdn,
		settings.TimeoutSeconds,
		nil,
	}
	// establishTrust generates the tlsConfig of the dnsClient; if not successful, the program will abort
	err := dnsClient.establishTrust()
	if err != nil {
		return nil, err
	}
	return &dnsClient, nil
}

// makeDnsProxyServer makes the dns proxy server and aborts if the tls server client fails its initial handshake
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
	proxyServer, err := settings.makeDnsProxyServer()
	if err != nil {
		return err
	}
	return proxyServer.listenAndServe()
}
