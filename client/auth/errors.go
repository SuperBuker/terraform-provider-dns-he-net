package auth

// ErrOTPDisabled is an error returned when the OTP is disabled.
type ErrOTPDisabled struct {
}

// Error returns a human-readable error message.
func (e *ErrOTPDisabled) Error() string {
	return "otp is not enabled"
}

// Unwrap implements the errors.Unwrap interface.
func (e *ErrOTPDisabled) Unwrap() []error {
	return []error{}
}

// ErrFileIO is an error returned when there is an issue reading from or
// writing to a file.
type ErrFileIO struct {
	err error
}

// Error returns a human-readable error message.
func (e *ErrFileIO) Error() string {
	return "file reading/writing failed"
}

// Unwrap returns the underlying error that caused this error.
func (e *ErrFileIO) Unwrap() []error {
	return []error{
		e.err,
	}
}

// ErrFileEncoding is an error returned when there is an issue decoding or
// encoding a file.
type ErrFileEncoding struct {
	err error
}

// Error returns a human-readable error message.
func (e *ErrFileEncoding) Error() string {
	return "file serialisation/deserialisation failed"
}

// Unwrap returns the underlying error that caused this error.
func (e *ErrFileEncoding) Unwrap() []error {
	return []error{
		e.err,
	}
}

// ErrFileChecksum is returned by the Reader when the file checksum fails to
// validate.
type ErrFileChecksum struct {
}

// Error returns a human-readable error message.
func (e *ErrFileChecksum) Error() string {
	return "file checksum validation failed"
}

// Unwrap implements the errors.Unwrap interface.
func (e *ErrFileChecksum) Unwrap() []error {
	return []error{}
}

// ErrFileEncrypt is returned by the when the file encryption fails.
type ErrFileEncrypt struct {
	err error
}

// Error returns a human-readable error message.
func (e *ErrFileEncrypt) Error() string {
	return "file encryption/decryption failed"
}

// Unwrap returns the underlying error that caused this error.
func (e *ErrFileEncrypt) Unwrap() []error {
	return []error{
		e.err,
	}
}
