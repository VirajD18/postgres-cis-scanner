package models

type Inventory struct {
	Host              string
	HAHosts		[]string
	DRHost		  string
	DatabaseName      string
	PostgresVersion   string
	Platform          string
	ManagedService    bool

	Benchmark         string

	Extensions        []string
	Modules           []string
}
