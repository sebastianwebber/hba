package lib

import "net"

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
