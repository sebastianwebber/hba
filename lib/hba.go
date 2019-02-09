package lib

import (
	"fmt"
	"net"
)

// HbaRule : Details a rule from the pg_hba.conf file
type HbaRule struct {
	Type         string
	DatabaseName string
	UserName     string
	DNSAddress   string
	IPAddress    net.IP
	NetworkMask  *net.IPMask
	AuthMethod   string
	LineNumber   int
	Comments     string
}

func (r *HbaRule) String() string {
	return formatRule(r)
}

func formatRule(rule *HbaRule) string {

	if rule == nil {
		return ""
	}

	if rule.Type == "local" {
		return formatLocal(*rule)
	}

	return ""
}

func formatLocal(r HbaRule) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", r.Type, r.DatabaseName, r.UserName, r.AuthMethod)
}
