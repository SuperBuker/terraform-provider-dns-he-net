// *Notice:*
// This module is mostly oriented to validate the DDNS credentials.
// This is mostly a "hack" to validate the current DDNS status.
//
// The functionalities related to update DDNS records are just for testing purposes.
// No support or enhancement will be provided for this module.
//
// For more information about the API, please refer to:
// https://dns.he.net/docs.html
//
// *About the API rate limit:*
// The API enfoces a rate limit based on the requested domain, thus it's only
// safe to fail one or two times the authentication against the same domain.
//
// This forces us to randomise the domains used during testing.
package ddns
