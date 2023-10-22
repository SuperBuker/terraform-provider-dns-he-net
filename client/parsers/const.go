package parsers

const (
	// accountQ is the XPath query for the account field.
	accountQ = "//form[@name='remove_domain']/input[@name='account']"

	// loginOkQ is the XPath query for the logout hyperlink.
	loginOkQ = "//a[@id='_tlogout']"

	// loginOtpQ is the XPath query for the OTP form.
	loginOtpQ = "//input[@id='tfacode']"

	// loginNoAuthQ is the XPath query for the login form.
	loginNoAuthQ = "//form[@name='login']"

	// zonesTableQ is the XPath query for the zones table.
	zonesTableQ = "//table[@id='domains_table']"

	// zoneQ is the XPath query for the zone rows.
	zoneQ = zonesTableQ + "/tbody/tr"

	// zoneIDQ is the XPath query for the zone ID within the zone row.
	zoneIDQ = "//td[@style]/img[@name][@value]"

	// recordsTableQ is the XPath query for the records table.
	recordsTableQ = "//div[@id='dns_main_content']/table[@class='generictable']"

	// recordQ is the XPath query for the record rows.
	recordQ = recordsTableQ + "/tbody/tr[@class]"

	// statusQ is the XPath query for the status message.
	statusQ = "//div[@id='dns_status']"

	// errorQ is the XPath query for the error message.
	errorQ = "//div[@id='dns_err']"
)
