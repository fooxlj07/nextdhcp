package serve

import (
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/nextdhcp/nextdhcp/pkg/middleware"
)

func prepareDHCPv4RequestReply(ctx *middleware.Context, req *dhcpv4.DHCPv4, s *SubnetConfig) (*dhcpv4.DHCPv4, error) {
	return nil, nil
}