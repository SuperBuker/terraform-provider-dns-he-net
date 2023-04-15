package params

// DomainCreate returnes the parameters necessary to create a tld domain.
func DomainCreate(m map[string]string) map[string]string {
	m["action"] = "add_zone"
	m["retmain:"] = "0"
	m["submit"] = "Add Domain!"

	return m
}

// DomainDelete returnes the parameters necessary to delete a tld domain.
func DomainDelete(m map[string]string) map[string]string {
	m["remove_domain"] = "1"

	return m
}
