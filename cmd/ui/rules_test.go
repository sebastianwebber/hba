package ui

import (
	"bufio"
	"bytes"
	"net"
	"reflect"
	"testing"

	"github.com/sebastianwebber/hba"
)

var (
	localRule    = hba.Rule{LineNumber: 1, Type: "local", DatabaseName: "all", UserName: "all", AuthMethod: "trust", Comments: "comment goes here"}
	localRuleStr = `
+------+-------+----------+------+-----------+--------+-------------------+
| LINE | TYPE  | DATABASE | USER | ADDRESSES | METHOD |     COMMENTS      |
+------+-------+----------+------+-----------+--------+-------------------+
|    1 | local | all      | all  |           | trust  | comment goes here |
+------+-------+----------+------+-----------+--------+-------------------+
`

	ip, netmask, _ = net.ParseCIDR("192.168.150.0/22")
	hostRuleCIDR   = hba.Rule{LineNumber: 2, Type: "host", DatabaseName: "all", UserName: "all", IPAddress: ip, NetworkMask: &netmask.Mask, AuthMethod: "trust", Comments: "comment goes here"}

	hostRuleDNS = hba.Rule{LineNumber: 3, Type: "host", DatabaseName: "all", UserName: "all", DNSAddress: "super-site.com", AuthMethod: "md5", Comments: "comment goes here"}

	hostRuleStr = `
+------+------+----------+------+------------------+--------+-------------------+
| LINE | TYPE | DATABASE | USER |    ADDRESSES     | METHOD |     COMMENTS      |
+------+------+----------+------+------------------+--------+-------------------+
|    2 | host | all      | all  | 192.168.150.0/22 | trust  | comment goes here |
|    3 | host | all      | all  | super-site.com   | md5    | comment goes here |
+------+------+----------+------+------------------+--------+-------------------+
`
	hostAllStr = `
+------+-------+----------+------+------------------+--------+-------------------+
| LINE | TYPE  | DATABASE | USER |    ADDRESSES     | METHOD |     COMMENTS      |
+------+-------+----------+------+------------------+--------+-------------------+
|    1 | local | all      | all  |                  | trust  | comment goes here |
|    2 | host  | all      | all  | 192.168.150.0/22 | trust  | comment goes here |
|    3 | host  | all      | all  | super-site.com   | md5    | comment goes here |
+------+-------+----------+------+------------------+--------+-------------------+
`
)

func makeRules(local, hostCIDR, hostDNS bool) []hba.Rule {

	out := []hba.Rule{}

	if local {
		out = append(out, localRule)
	}

	if hostCIDR {
		out = append(out, hostRuleCIDR)
	}

	if hostDNS {
		out = append(out, hostRuleDNS)
	}

	return out
}

func TestDisplayRules(t *testing.T) {

	tests := []struct {
		name string
		args []hba.Rule
		want string
	}{
		{name: "should display a nice message when rules are empty", want: "No rules found\n"},
		{name: "should display a nice table for local rules", args: makeRules(true, false, false), want: localRuleStr},
		{name: "should display a nice table for both host rules", args: makeRules(false, true, true), want: hostRuleStr},
		{name: "should display a nice table for all rules", args: makeRules(true, true, true), want: hostAllStr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var buf bytes.Buffer
			writer := bufio.NewWriter(&buf)

			DisplayRules(tt.args, writer)

			got := buf.String()

			if len(tt.args) > 0 {
				got = "\n" + got
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisplayRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
