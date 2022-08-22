package main

type Time int

// START OMIT
// MarshalJSON implements the json.Marshaler interface. // HL
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t Time) MarshalJSON() ([]byte, error) { // HL
	// Redacted
	return nil, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. // HL
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error { // HL
	// Redacted
	return nil
}

// END OMIT
