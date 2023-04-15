package auth

// Status indicates the current authentication status.
type Status int8

const (
	NoAuth Status = iota
	Ok
	OTP
	Unknown
	Other
)
