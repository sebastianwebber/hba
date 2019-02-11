package hba

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

const (
	regexDBUser   = `(?P<dbname>[\w\,]+)\s+(?P<username>[\w\,]+)`
	regexComments = `(?:(?:[ \t]+)?(?:[#]+)?(?:[ \t]+)?(?P<comment>[\w \t\:]+))?`
	regexMethod   = `(?P<method>\w+)`
	regexHost     = `(?P<type>host(?:no)?(?:ssl)?)\s+` + regexDBUser + `\s+(?:(?P<address>[\w\.\/\:\-]+)(?:\s+(?P<mask>[\d\.]+))?)\s+` + regexMethod + regexComments
	regexLocal    = `(?P<type>local)\s+` + regexDBUser + `\s+` + regexMethod + regexComments
)

// Parse a line from pg_hba.conf file
func Parse(line string) (*HbaRule, error) {
	return parseLine(line)
}

func parseLine(line string) (*HbaRule, error) {

	if strings.HasPrefix(line, "local") {
		return parseLocal(line)
	}

	return parseHost(line)
}

var reLocal = regexp.MustCompile(regexLocal)

func parseLocal(line string) (*HbaRule, error) {

	matches := reLocal.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return nil, errors.New("invalid line")
	}

	return parseLocalParts(matches[0][1:]), nil
}

var reHost = regexp.MustCompile(regexHost)

func parseHost(line string) (*HbaRule, error) {

	matches := reHost.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return nil, errors.New("invalid line")
	}

	return parseHostParts(matches[0][1:]), nil
}

func parseLocalParts(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		AuthMethod:   parts[3],
		Comments:     parts[4],
	}
}

func parseHostParts(parts []string) *HbaRule {

	if parts[4] == "" {
		if strings.Contains(parts[3], "/") {
			return parseHostOctet(parts)
		}

		return parseHostDNS(parts)
	}

	mask := net.IPMask(net.ParseIP(parts[4]))

	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		IPAddress:    net.ParseIP(parts[3]),
		NetworkMask:  &mask,
		AuthMethod:   parts[5],
		Comments:     parts[6],
	}
}

func parseHostOctet(parts []string) *HbaRule {

	addr, mask, _ := net.ParseCIDR(parts[3])

	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		IPAddress:    addr,
		NetworkMask:  &mask.Mask,
		AuthMethod:   parts[5],
		Comments:     parts[6],
	}
}

func parseHostDNS(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		DNSAddress:   parts[3],
		AuthMethod:   parts[5],
		Comments:     parts[6],
	}
}
