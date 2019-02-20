package hba

import "testing"

func TestRules_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		c       *Rules
		args    Rule
		wantErr bool
	}{
		{name: "should throw an eror when rule is invalid", c: &Rules{}, args: Rule{}, wantErr: true},
		{name: "should add a valid rule without error", c: &Rules{}, args: localRule},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.Add(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rules.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			totalRules := len(*tt.c)
			if totalRules == 0 && (err == nil) {
				t.Errorf("Expected 1 rule, got %d", totalRules)
			}
		})
	}
}
