# PostgreSQL CIS Scanner

A Go-based PostgreSQL CIS Benchmark security scanner for auditing PostgreSQL servers against security controls.

## Features

- PostgreSQL CIS Benchmark scanning
- PostgreSQL 18 support
- Multiple PostgreSQL server scanning
- IaaS and PASS environment support
- Custom CIS control templates
- HTML reports
- JSON results
- HA and DR server configuration
- RHEL 9 RPM package

## Installation

Install the RPM:

```bash
rpm -ivh pgcis-1.0.0-1.el9.x86_64.rpm
```

Verify installation:

```bash
rpm -q pgcis
```

## Configuration

The main configuration file is:

`/etc/pgcis/servers.json`

Example:

```json
[
  {
    "name": "Production PostgreSQL",
    "host": "postgres.example.com",
    "port": 5432,
    "database": "postgres",
    "user": "pgcis_scanner",
    "password": "CHANGE_ME",
    "sslmode": "require",
    "type": "iaas",
    "control_template": "/etc/pgcis/templates/iaas.json",
    "ha_hosts": [],
    "dr_host": ""
  }
]
```

## Server Types

### IaaS

```json
"type": "iaas",
"control_template": "/etc/pgcis/templates/iaas.json"
```

### PASS

```json
"type": "pass",
"control_template": "/etc/pgcis/templates/pass.json"
```

## Control Templates

Templates define which CIS controls should be scanned.

IaaS template:

`/etc/pgcis/templates/iaas.json`

PASS template:

`/etc/pgcis/templates/pass.json`

Example:

```json
{
  "controls": [
    "3.1.2",
    "3.1.3",
    "3.1.20",
    "6.8",
    "6.8.1"
  ]
}
```

Only the Control IDs listed in the template are scanned.

### Custom Template

Create your own template, for example:

`/etc/pgcis/templates/custom.json`

Example:

```json
{
  "controls": [
    "3.1.2",
    "6.8",
    "7.4"
  ]
}
```

Reference it from `servers.json`:

```json
"control_template": "/etc/pgcis/templates/custom.json"
```

If `control_template` is not specified, all available CIS controls are scanned.

## Running the Scanner

Run:

```bash
pgcis
```

The scanner will:

1. Load configured PostgreSQL servers.
2. Connect to each server.
3. Detect the PostgreSQL version.
4. Load the appropriate CIS benchmark.
5. Apply the configured control template.
6. Execute the selected controls.
7. Generate reports.

## Reports

Reports are generated separately for each server.

Report location:

`/var/lib/pgcis/reports/`

Each server has its own report directory containing the generated reports.

## Development

Build:

```bash
go build ./...
```

Run directly:

```bash
go run ./cmd/pgcis
```

## Security

Do not commit real PostgreSQL passwords, credentials, API keys, certificates, or production connection details to GitHub.

Use `CHANGE_ME` for example credentials.

## Project Structure

```text
postgres-cis-scanner/
├── benchmark/
│   └── PostgreSQL18/
├── configs/
│   ├── templates/
│   │   ├── iaas.json
│   │   └── pass.json
│   └── README.txt
├── internal/
│   ├── benchmark/
│   ├── charts/
│   ├── database/
│   ├── engine/
│   ├── inventory/
│   ├── models/
│   ├── report/
│   ├── scanner/
│   └── servers/
├── templates/
├── go.mod
├── go.sum
└── README.md
```

## License

Proprietary
