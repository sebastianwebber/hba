package hba

import (
	"net"
	"reflect"
	"testing"
)

func parseIPMask(s string) *net.IPMask {
	mask := net.IPMask(net.ParseIP(s))
	return &mask
}

func Test_Parse(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    Rule
		wantErr bool
	}{
		{
			name:    "should throw an error on invalid lines - local",
			args:    "local#############",
			want:    Rule{},
			wantErr: true,
		},
		{
			name:    "should throw an error on invalid lines - host",
			args:    "host all md4",
			want:    Rule{},
			wantErr: true,
		},
		{
			name:    "should parse a local line",
			args:    localRuleStr,
			want:    localRule,
			wantErr: false,
		},
		{
			name:    "should parse a host line (ip/octet)",
			args:    hostRuleCIDRStr,
			want:    hostRuleCIDR,
			wantErr: false,
		},
		{
			name:    "should parse a host line (ip regular_mask)",
			args:    "host    all             all             192.168.150.0 255.255.252.0            trust",
			want:    Rule{Type: "host", DatabaseName: "all", UserName: "all", IPAddress: ip, NetworkMask: parseIPMask("255.255.252.0"), AuthMethod: "trust"},
			wantErr: false,
		},
		{
			name:    "should parse a host line (dns addressses)",
			args:    hostRuleDNSStr,
			want:    hostRuleDNS,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				if !reflect.DeepEqual(*got, tt.want) {
					t.Errorf("parseLine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
