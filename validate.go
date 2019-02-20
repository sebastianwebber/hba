package hba

import (
	"errors"
	"net"
	"strings"
)

// IsValid validates a Rule
func (r *Rule) IsValid() (bool, error) {
	return isValid(r)
}

var emptyIP net.IP

func isValid(r *Rule) (bool, error) {

	if !sliceExists(r.Type, allowedConnTypes) {
		return false, errors.New("Invalid connection type")
	}

	if !validateNonEmpty(r.DatabaseName) {
		return false, errors.New("Invalid database name")
	}

	if !validateNonEmpty(r.UserName) {
		return false, errors.New("Invalid username")
	}
	if !sliceExists(r.AuthMethod, allowedAuthMethods) {
		return false, errors.New("Invalid authentication method")
	}

	if strings.HasPrefix(r.Type, "host") {
		return validateHost(r)
	}

	return true, nil
}

var (
	allowedConnTypes = []string{
		"local",
		"host",
		"hostssl",
		"hostnossl",
	}
	allowedAuthMethods = []string{
		"md5",
		"trust",
		"reject",
		"password",
		"gss",
		"sspi",
		"krb5",
		"ident",
		"peer",
		"pam",
		"ldap",
		"radius",
		"cert",
	}
)

func validateNonEmpty(field string) bool {
	if field != "" {
		return true
	}

	return false
}

func sliceExists(val string, list []string) bool {

	for i := 0; i < len(list); i++ {
		if val == list[i] {
			return true
		}
	}

	return false
}

func validateHost(r *Rule) (bool, error) {

	if len(r.IPAddress) == 0 && !validateNonEmpty(r.DNSAddress) {
		return false, errors.New("Missing network or DNS address")
	}

	return true, nil
}
