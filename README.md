## Server Types

The `type` field in `servers.json` defines the PostgreSQL deployment platform.

Supported values:

| Type | Platform | Managed Service |
|------|----------|-----------------|
| `iaas` | `self-managed` | No |
| `rds` | `rds` | Yes |
| `aurora` | `aurora` | Yes |
| `azure-flex` | `azure-flex` | Yes |

### `iaas`

Use for PostgreSQL running on your own infrastructure, such as:

- RHEL/Linux servers
- AWS EC2
- Azure VMs
- VMware
- Bare-metal servers

Example:

```json
{
  "name": "Production PostgreSQL",
  "host": "postgres.example.com",
  "port": 5432,
  "database": "postgres",
  "user": "postgres",
  "password": "CHANGE_ME",
  "sslmode": "require",
  "type": "iaas",
  "control_template": "configs/templates/iaas.json",
  "ha_hosts": [],
  "dr_host": ""
}
```

### `rds`

Use for Amazon RDS for PostgreSQL.

```json
"type": "rds"
```

### `aurora`

Use for Amazon Aurora PostgreSQL.

```json
"type": "aurora"
```

Example:

```json
{
  "name": "Production Aurora",
  "host": "aurora-endpoint.amazonaws.com",
  "port": 5432,
  "database": "postgres",
  "user": "postgres",
  "password": "CHANGE_ME",
  "sslmode": "require",
  "type": "aurora",
  "control_template": "configs/templates/pass.json",
  "ha_hosts": [
    "aurora-read-replica.amazonaws.com"
  ],
  "dr_host": "aurora-dr-endpoint.amazonaws.com"
}
```

### `azure-flex`

Use for Azure Database for PostgreSQL Flexible Server.

```json
"type": "azure-flex"
```

### Control Templates

The `control_template` field determines which set of CIS controls is applied to the server.

Available templates:

```text
configs/templates/iaas.json
configs/templates/pass.json
```

For RPM installations, these are automatically resolved from:

```text
/etc/pgcis/templates/
```

You can keep the same `control_template` value in `servers.json` for both source and RPM installations.

### Important

Use the correct `type` because it affects:

- Platform shown in the report
- Managed Service status
- Platform-specific CIS controls
- Control applicability

For example, an Aurora PostgreSQL server should use:

```json
"type": "aurora"
```

and not:

```json
"type": "iaas"
```

or:

```json
"type": "pass"
```
