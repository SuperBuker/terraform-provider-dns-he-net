package params

// DomainCreate returns the parameters necessary to create a domain.
func DomainCreate(m map[string]string) map[string]string {
	m["action"] = "add_domain"
	m["retmain:"] = "0"
	m["submit"] = "Add Domain!"

	return m
}

// DomainDelete returns the parameters necessary to delete a domain.
func DomainDelete(m map[string]string) map[string]string {
	m["remove_domain"] = "1"

	return m
}
