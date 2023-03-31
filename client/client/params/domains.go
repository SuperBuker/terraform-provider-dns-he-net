package params

func DomainCreate(m map[string]string) map[string]string {
	m["action"] = "add_zone"
	m["retmain:"] = "0"
	m["submit"] = "Add Domain!"

	return m
}

func DomainDelete(m map[string]string) map[string]string {
	m["remove_domain"] = "1"

	return m
}
