package hba

import (
	"io"
	"strings"
	"testing"

	messagediff "gopkg.in/d4l3k/messagediff.v1"
)

func makeFile(coments bool, validLocal bool, validCIDR bool, validDNS bool, invalid bool) io.Reader {

	out := ""

	if coments {
		out += `
# comented line
##### commented line multi
  # comentt # coment again ##### end`
	}

	if validLocal {
		out += "\n" + localRuleStr
	}

	if validCIDR {
		out += "\n" + hostRuleCIDRStr
	}

	if validDNS {
		out += "\n" + hostRuleDNSStr
	}

	if invalid {
		out += "\n" + "local	trust"
	}

	return strings.NewReader(out)
}

func makeRules(validLocal, validCIDR, validDNS bool) *[]Rule {

	out := []Rule{}
	const comentsTotal = 4

	if validLocal {
		out = append(out, localRule)
		out[0].LineNumber = 1 + comentsTotal
	}

	if validCIDR {
		out = append(out, hostRuleCIDR)
		out[1].LineNumber = 2 + comentsTotal
	}

	if validDNS {
		out = append(out, hostRuleDNS)
		out[2].LineNumber = 3 + comentsTotal
	}

	return &out
}

func compare(t *testing.T, generated interface{}, expected interface{}) {
	if diff, equal := messagediff.PrettyDiff(generated, expected); !equal {
		t.Errorf("Generated = %#v", generated)
		t.Errorf("Expected = %#v\n", expected)
		t.Errorf("Diff = %s", diff)
	}
}

func TestParseReader(t *testing.T) {

	tests := []struct {
		name    string
		args    io.Reader
		want    *[]Rule
		wantErr bool
	}{
		{name: "should throw an error when buffer is nil", wantErr: true},
		{name: "should ignore lines starting with #", args: makeFile(true, true, false, false, false), want: makeRules(true, false, false)},
		{name: "should handle invalid lines", args: makeFile(false, false, false, false, true), wantErr: true},
		{name: "should parse a complete (and valid) file", args: makeFile(true, true, true, true, false), want: makeRules(true, true, true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReader(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compare(t, got, tt.want)
		})
	}
}
