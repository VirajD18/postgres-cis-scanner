PGCIS SERVER CONFIGURATION
==========================

The server configuration is stored in:

    /etc/pgcis/servers.json


1. BASIC SERVER CONFIGURATION
-----------------------------

Each PostgreSQL server is added as an object inside servers.json.

Example:

{
  "name": "Production PostgreSQL",
  "host": "localhost",
  "port": 5432,
  "database": "postgres",
  "user": "postgres",
  "password": "Change_me",
  "sslmode": "require",
  "type": "iaas",
  "control_template": "/etc/pgcis/templates/iaas.json",
  "ha_hosts": [],
  "dr_host": ""
}


2. SERVER TYPE
--------------

The "type" identifies the server environment.

IaaS:

    "type": "iaas"

PASS:

    "type": "pass"


3. CONTROL TEMPLATE
--------------------

The "control_template" specifies which CIS controls should
be executed for that server.

IaaS example:

    "type": "iaas",
    "control_template": "/etc/pgcis/templates/iaas.json"

PASS example:

    "type": "pass",
    "control_template": "/etc/pgcis/templates/pass.json"


4. TEMPLATE FORMAT
-------------------

A template contains the Control IDs that should be scanned.

Example:

{
  "controls": [
    "3.1.2",
    "3.1.3",
    "3.1.20",
    "6.8",
    "6.8.1"
  ]
}


5. CUSTOM TEMPLATES
-------------------

You can create your own template.

Example:

    /etc/pgcis/templates/custom.json

Then add:

    "control_template": "/etc/pgcis/templates/custom.json"

Only the Control IDs listed in that template will be scanned.


6. NO TEMPLATE
--------------

If "control_template" is empty or omitted:

    "control_template": ""

the scanner will execute all available CIS controls.


7. MULTIPLE SERVERS
-------------------

Multiple servers can be configured in the same servers.json.

Each server can use a different control template.

Example:

Server A:
    type = iaas
    template = iaas.json

Server B:
    type = pass
    template = pass.json

Server C:
    no template
    → all CIS controls


8. CONTROL IDs
--------------

Control IDs must match the IDs available in the installed
PostgreSQL CIS benchmark.

Example:

    3.1.2
    3.1.3
    6.8
    6.8.1


9. SECURITY
-----------

Do not commit servers.json containing real passwords,
credentials, or production connection details to GitHub.

Use a sanitized example configuration for source control.
