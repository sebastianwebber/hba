package lib

// HbaRule : Details a rule from the pg_hba.conf file
type HbaRule struct {
	Type         string
	DatabaseName string
	UserName     string
	DNSAddress   string
	IPAddress    string
	NetworkMask  string
	AuthMethod   string
	LineNumber   int
	Comments     string
}
