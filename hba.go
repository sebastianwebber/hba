package hba

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

	return formatHost(*rule)
}

func formatLocal(r HbaRule) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t# %s", r.Type, r.DatabaseName, r.UserName, r.AuthMethod, r.Comments)
}

func formatHost(r HbaRule) string {

	if r.DNSAddress != "" {
		return fmt.Sprintf(
			"%s\t%s\t%s\t%s\t%s\t# %s",
			r.Type,
			r.DatabaseName,
			r.UserName,
			r.DNSAddress,
			r.AuthMethod,
			r.Comments,
		)
	}

	octMask, _ := r.NetworkMask.Size()

	return fmt.Sprintf(
		"%s\t%s\t%s\t%s/%d\t%s\t# %s",
		r.Type,
		r.DatabaseName,
		r.UserName,
		r.IPAddress.String(),
		octMask,
		r.AuthMethod,
		r.Comments,
	)
}
