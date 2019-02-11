package lib

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

const (
	regexDBUser   = `(?P<dbname>[\w\,]+)\s+(?P<username>[\w\,]+)`
	regexComments = `(?:(?:[ \t]+)?(?:[#]+)?(?:[ \t]+)?(?P<comment>[\w \t\:]+))`
	regexMethod   = `(?P<method>\w+)`
	regexHost     = `(?P<type>host(?:no)?(?:ssl)?)\s+` + regexDBUser + `\s+(?:(?P<address>[\w\.\/\:\-]+)(?:\s+(?P<mask>[\d\.]+))?)\s+` + regexMethod
	regexLocal    = `(?P<type>local)\s+` + regexDBUser + `\s+` + regexMethod + regexComments
)

// Parse a line from pg_hba.conf file
func Parse(line string) (*HbaRule, error) {
	return parseLine(line)
}

func parseLine(line string) (*HbaRule, error) {

	parts := strings.Fields(line)

	var err error
	if err = validateLine(parts); err != nil {
		return nil, fmt.Errorf("could not parse line: %v", err)
	}

	if strings.HasPrefix(line, "local") {
		return parseLocal(parts), nil
	}

	return parseHost(line)
}

var reHost = regexp.MustCompile(regexHost)

func parseHost(line string) (*HbaRule, error) {

	matches := reHost.FindAllStringSubmatch(line, -1)

	return parseHostParts(matches[0][1:]), nil
}

func parseLocal(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		AuthMethod:   parts[3],
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
	}
}

func parseHostDNS(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		DNSAddress:   parts[3],
		AuthMethod:   parts[5],
	}
}

func validateLine(parts []string) error {
	if len(parts) < 4 {
		return errors.New("invalid fields length")
	}

	return nil
}
