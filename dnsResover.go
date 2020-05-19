package resolver

import (
	"fmt"
	"github.com/BASChain/go-bmail-account"
	"net"
)

type DNSResolver struct {
}

func (ds *DNSResolver) DomainA(domain string) []net.IP {
	fmt.Println("implement me")
	return nil
}

func (ds *DNSResolver) DomainMX(domainMX string) ([]net.IP, []bmail.Address) {
	fmt.Println("implement me")
	return nil, nil
}

func (ds *DNSResolver) BMailBCA(mailHash string) (bmail.Address, string) {
	fmt.Println("implement me")
	return "", ""
}

func NewDnsResolver() NameResolver {
	obj := &DNSResolver{}

	return obj
}
