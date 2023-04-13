package auth

type ErrFileIO struct {
	err error
}

func (e *ErrFileIO) Error() string {
	return "file reading/writing failed"
}

func (e *ErrFileIO) Unwrap() []error {
	return []error{
		e.err,
	}
}

type ErrFileEncodng struct {
	err error
}

func (e *ErrFileEncodng) Error() string {
	return "file serialisation/deserialisation failed"
}

func (e *ErrFileEncodng) Unwrap() []error {
	return []error{
		e.err,
	}
}

type ErrFileChecksum struct {
}

func (e *ErrFileChecksum) Error() string {
	return "file checksum validation failed"
}

func (e *ErrFileChecksum) Unwrap() []error {
	return []error{}
}

type ErrFileEncrypt struct {
	err error
}

func (e *ErrFileEncrypt) Error() string {
	return "file encryption/decryption failed"
}

func (e *ErrFileEncrypt) Unwrap() []error {
	return []error{
		e.err,
	}
}
