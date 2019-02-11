package lib

import (
	"net"
	"testing"
)

var (
	localRule    = HbaRule{Type: "local", DatabaseName: "all", UserName: "all", AuthMethod: "trust", Comments: "comment goes here"}
	localRuleStr = "local	all	all	trust	# comment goes here"

	ip, netmask, _  = net.ParseCIDR("192.168.150.0/22")
	hostRuleCIDR    = HbaRule{Type: "host", DatabaseName: "all", UserName: "all", IPAddress: ip, NetworkMask: &netmask.Mask, AuthMethod: "trust", Comments: "comment goes here"}
	hostRuleCIDRStr = "host	all	all	192.168.150.0/22	trust	# comment goes here"

	hostRuleDNS    = HbaRule{Type: "host", DatabaseName: "all", UserName: "all", DNSAddress: "super-site.com", AuthMethod: "trust", Comments: "comment goes here"}
	hostRuleDNSStr = "host	all	all	super-site.com	trust	# comment goes here"
)

func Test_formatRule(t *testing.T) {
	tests := []struct {
		name string
		args *HbaRule
		want string
	}{
		{name: "should return empty when rule is nil", want: ""},
		{name: "should parse a local rule", args: &localRule, want: localRuleStr},
		{name: "should parse a host rule with CIDR addresses", args: &hostRuleCIDR, want: hostRuleCIDRStr},
		{name: "should parse a host rule with DNS addresses", args: &hostRuleDNS, want: hostRuleDNSStr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.String(); got != tt.want {
				t.Errorf("formatRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
