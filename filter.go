package hba

import (
	"strings"
)

// Filter creates a new collection of []Rule based on the input string
func Filter(rules []Rule, filter string) []Rule {

	if filter == "" {
		return rules
	}

	var out []Rule

	for i := 0; i < len(rules); i++ {
		ruleStr := rules[i].String()

		if strings.Contains(ruleStr, filter) {
			out = append(out, rules[i])
		}
	}

	return out
}
