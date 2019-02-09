package lib

import (
	"net"
	"reflect"
	"testing"
)

func parseIPMask(s string) *net.IPMask {

	mask := net.IPMask(net.ParseIP(s))
	return &mask
}

func Test_parseLine(t *testing.T) {

	ip, netmask, _ := net.ParseCIDR("127.0.0.1/32")

	tests := []struct {
		name    string
		args    string
		want    HbaRule
		wantErr bool
	}{
		{
			name:    "should parse a local line",
			args:    localRuleStr,
			want:    localRule,
			wantErr: false,
		},
		{
			name:    "should parse a host line (ip/octet)",
			args:    "host    all             all             127.0.0.1/32            trust",
			want:    HbaRule{Type: "host", DatabaseName: "all", UserName: "all", IPAddress: ip, NetworkMask: &netmask.Mask, AuthMethod: "trust"},
			wantErr: false,
		},
		{
			name:    "should parse a host line (ip regular_mask)",
			args:    "host    all             all             127.0.0.1 255.255.255.0            trust",
			want:    HbaRule{Type: "host", DatabaseName: "all", UserName: "all", IPAddress: ip, NetworkMask: parseIPMask("255.255.255.0"), AuthMethod: "trust"},
			wantErr: false,
		},
		{
			name:    "should parse a host line (dns addressses)",
			args:    "host    all             all             super-site.com            trust",
			want:    HbaRule{Type: "host", DatabaseName: "all", UserName: "all", DNSAddress: "super-site.com", AuthMethod: "trust"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLine(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
