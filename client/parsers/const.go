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

	// tabsStandardQ is the XPath query for the standard tabs div.
	tabsStandardQ = "//div[@id='tabs-standard']"

	// domainZonesTableQ is the XPath query for the domains table.
	domainZonesTableQ = tabsStandardQ + "//table[@id='domains_table' and contains(@class, 'generic_table')]"

	// domainZoneQ is the XPath query for the domain rows.
	domainZoneQ = domainZonesTableQ + "/tbody/tr"

	// domainZoneIDQ is the XPath query for the domain ID within the domain row.
	domainZoneIDQ = "//td[@style]/img[@name][@value]"

	// prefixesTableQ is the XPath query for the delegated prefixes table.
	prefixesTableQ = tabsStandardQ + "//table[not(@id) and @class='generic_table']"

	// prefixQ is the XPath query for the record rows.
	prefixQ = prefixesTableQ + "/tbody/tr"

	// prefixIDQ is the XPath query for the prefix ID within the prefix row.
	prefixIDQ = "//td[@style]/img[@value]"

	// prefixNameQ is the XPath query for the prefix ID within the prefix row.
	prefixNameQ = "//td[@class='delegated']"

	// tabsAdvancedQ is the XPath query for the standard tabs div.
	tabsAdvancedQ = "//div[@id='tabs-advanced']"

	// arpaZoneTableQ is the XPath query for the delegated ARPA zones table.
	arpaZoneTableQ = tabsAdvancedQ + "//table[not(@id) and @class='generic_table']"

	// arpaZoneQ is the XPath query for the record rows.
	arpaZoneQ = arpaZoneTableQ + "/tbody/tr"

	// arpaZoneIDQ is the XPath query for the ARPA zone ID within the ARPA zone row.
	arpaZoneIDQ = "//td[@style]/img[@name][@value]"

	// recordsTableQ is the XPath query for the records table.
	recordsTableQ = "//div[@id='dns_main_content']/table[@class='generictable']"

	// recordQ is the XPath query for the record rows.
	recordQ = recordsTableQ + "/tbody/tr[@class]"

	// statusQ is the XPath query for the status message.
	statusQ = "//div[@id='dns_status']"

	// errorQ is the XPath query for the error message.
	errorQ = "//div[@id='dns_err']"
)
