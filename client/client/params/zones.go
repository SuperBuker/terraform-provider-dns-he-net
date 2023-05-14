package params

// ZoneCreate returnes the parameters necessary to create a zone.
func ZoneCreate(m map[string]string) map[string]string {
	m["action"] = "add_zone"
	m["retmain:"] = "0"
	m["submit"] = "Add Domain!"

	return m
}

// ZoneDelete returnes the parameters necessary to delete a zone.
func ZoneDelete(m map[string]string) map[string]string {
	m["remove_domain"] = "1"

	return m
}
