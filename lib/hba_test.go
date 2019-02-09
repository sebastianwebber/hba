package lib

import (
	"testing"
)

var (
	localRule    = HbaRule{Type: "local", DatabaseName: "all", UserName: "all", AuthMethod: "trust"}
	localRuleStr = "local	all	all	trust"
)

func Test_formatRule(t *testing.T) {
	tests := []struct {
		name string
		args *HbaRule
		want string
	}{
		{name: "should return empty when rule is nil", want: ""},
		{name: "should parse a local rule", args: &localRule, want: localRuleStr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.String(); got != tt.want {
				t.Errorf("formatRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
