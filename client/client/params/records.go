package params

// recordGeneric returnes the parameters common to all record operations.
func recordGeneric(m map[string]string) map[string]string {
	m["menu"] = "edit_zone"
	m["hosted_dns_editzone"] = "1"

	return m
}

// RecordCreate returnes the parameters necessary to create a record.
func RecordCreate(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_editrecord"] = "Submit"

	return m
}

// RecordUpdate returnes the parameters necessary to update a record.
func RecordUpdate(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_editrecord"] = "Update"

	return m
}

// RecordDelete returnes the parameters necessary to delete a record.
func RecordDelete(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_delconfirm"] = "delete"
	m["hosted_dns_delrecord"] = "1"

	return m
}
