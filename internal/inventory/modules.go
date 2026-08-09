package inventory

func detectModules(platform string) []string {

	switch platform {

	case "Self Managed PostgreSQL":
		return []string{
			"postgres",
			"linux",
		}

	case "AWS Aurora PostgreSQL":
		return []string{
			"postgres",
			"aurora",
		}

	case "Azure PostgreSQL Flexible Server":
		return []string{
			"postgres",
			"azure",
		}

	case "Amazon RDS PostgreSQL":
		return []string{
			"postgres",
			"rds",
		}

	default:
		return []string{
			"postgres",
		}
	}
}
