package params

// DDNSKeySet returnes the parameters necessary to set the ddns key.
func DDNSKeySet(m map[string]string) map[string]string {
	recordGeneric(m)
	m["generate_key"] = "Submit"

	return m
}
