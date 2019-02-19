package ui

import (
	"reflect"
	"testing"

	"github.com/sebastianwebber/hba"
)

var (
	allRows = []hba.Rule{
		localRule,
		hostRuleCIDR,
		hostRuleDNS,
	}
)

func makeRows(local, host, md5Only bool) []hba.Rule {
	var out []hba.Rule

	if local {
		out = append(out, localRule)
	}

	if host {
		out = append(out, hostRuleCIDR, hostRuleDNS)
	}

	if md5Only {
		out = append(out, hostRuleDNS)
	}

	return out
}

func TestFilter(t *testing.T) {
	type args struct {
		rules  []hba.Rule
		filter string
	}
	tests := []struct {
		name string
		args args
		want []hba.Rule
	}{
		{name: "should ignore empty filter", want: allRows, args: args{rules: allRows, filter: ""}},
		{name: "should bring only local rows", want: makeRows(true, false, false), args: args{rules: allRows, filter: "local"}},
		{name: "should bring only host rows", want: makeRows(false, true, false), args: args{rules: allRows, filter: "host"}},
		{name: "should bring only rows with md5 auth", want: makeRows(false, false, true), args: args{rules: allRows, filter: "md5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.rules, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
