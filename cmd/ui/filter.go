package ui

import (
	"strings"

	"github.com/sebastianwebber/hba"
)

// Filter creates a new collection of []hba.Rule based on the input string
func Filter(rules []hba.Rule, filter string) []hba.Rule {

	if filter == "" {
		return rules
	}

	var out []hba.Rule

	for i := 0; i < len(rules); i++ {
		ruleStr := rules[i].String()

		if strings.Contains(ruleStr, filter) {
			out = append(out, rules[i])
		}
	}

	return out
}
