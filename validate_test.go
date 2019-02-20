package hba

import (
	"testing"
)

func TestRule_IsValid(t *testing.T) {

	tests := []struct {
		name    string
		fields  Rule
		want    bool
		wantErr bool
	}{
		{name: "should throw an error when a wrong connection type is set", fields: Rule{Type: "fooo"}, wantErr: true},
		{name: "should throw an error when database is missing", fields: Rule{Type: "local"}, wantErr: true},
		{name: "should throw an error user is missing", fields: Rule{Type: "local", DatabaseName: "sameuser"}, wantErr: true},
		{name: "should throw an error when a wrong auth method is set", fields: Rule{Type: "local", DatabaseName: "sameuser", UserName: "postgres", AuthMethod: "foobar"}, wantErr: true},
		{name: "should run ok with a valid local rule", fields: localRule, want: true, wantErr: false},
		{name: "on host conn type should throw an error when a dns or ip addresses is missing", fields: Rule{Type: "host", DatabaseName: "sameuser", UserName: "postgres", AuthMethod: "md5"}, want: false, wantErr: true},
		{name: "should run ok with a valid host CIDR rule", fields: hostRuleCIDR, want: true, wantErr: false},
		{name: "should run ok with a valid host DNS rule", fields: hostRuleDNS, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				Type:         tt.fields.Type,
				DatabaseName: tt.fields.DatabaseName,
				UserName:     tt.fields.UserName,
				DNSAddress:   tt.fields.DNSAddress,
				IPAddress:    tt.fields.IPAddress,
				NetworkMask:  tt.fields.NetworkMask,
				AuthMethod:   tt.fields.AuthMethod,
				LineNumber:   tt.fields.LineNumber,
				Comments:     tt.fields.Comments,
			}
			got, err := r.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rule.IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compare(t, got, tt.want)
		})
	}
}
