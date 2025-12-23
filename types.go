package enomapi

type Domain struct {
	Name      string
	Extension string
}

type DomainSetHostRequest struct {
	Domain Domain
	Hosts  []DomainSetHost
}

type DomainSetHost struct {
	HostName   string
	RecordType string
	Address    string
	MXPref     int
}
