package lib

// HbaRule : Details a rule from the pg_hba.conf file
type HbaRule struct {
	ConnectionType string
	DatabaseName   string
	UserName       string
	IPAddress      string
	NetworkMask    string
	AuthType       string
	LineNumber     int
	Comments       string
}
