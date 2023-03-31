package auth

type Status int8

const (
	NoAuth Status = iota
	Ok
	OTP
	Unknown
	Other
)
