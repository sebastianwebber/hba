package lib

import (
	"errors"
	"fmt"
	"net"
	"strings"
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
	ruleType := strings.ToLower(parts[0])

	switch ruleType {
	case "local":
		return parseLocal(parts), nil
	case "host":
		return parseHost(parts), nil
	}

	return nil, nil
}

func parseLocal(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		AuthMethod:   parts[3],
	}
}

func parseHost(parts []string) *HbaRule {

	if len(parts) == 5 {
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
		AuthMethod:   parts[4],
	}
}

func parseHostDNS(parts []string) *HbaRule {
	return &HbaRule{
		Type:         parts[0],
		DatabaseName: parts[1],
		UserName:     parts[2],
		DNSAddress:   parts[3],
		AuthMethod:   parts[4],
	}
}

func validateLine(parts []string) error {
	if len(parts) < 4 {
		return errors.New("invalid fields length")
	}

	return nil
}
