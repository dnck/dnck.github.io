package dnsserver

import (
  "fmt"
  "github.com/miekg/dns"
)

func (settings *Settings) MakeDnsProxyServer() error {
  dnsClient := client{settings.DnsResolverPin,
    "1.1.1.1:853",
    "cloudflare-dns.com",
  }
  err := dnsClient.establishTrust()
  if err != nil {
    fmt.Println(err.Error())
  }

  tcpServer := &dns.Server{Addr: settings.formatAddress(), Net: "tcp"}

  dns.HandleFunc(".", dnsClient.request)

  if err := tcpServer.ListenAndServe(); err != nil {
    return err
  }
  return nil
}

