package params

func recordGeneric(m map[string]string) map[string]string {
	m["menu"] = "edit_zone"
	m["hosted_dns_editzone"] = "1"

	return m
}

func RecordCreate(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_editrecord"] = "Submit"

	return m
}

func RecordUpdate(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_editrecord"] = "Update"

	return m
}

func RecordDelete(m map[string]string) map[string]string {
	recordGeneric(m)
	m["hosted_dns_delconfirm"] = "delete"
	m["hosted_dns_delrecord"] = "1"

	return m
}
