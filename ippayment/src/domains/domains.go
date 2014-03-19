package domains

import (
	
"github.com/mitchellh/mapstructure"
)

type Hits struct {
	
	Created int64
	Id	string
	Msisdn string
	Site   string
	Themes   string
	Resource string
}

type DecoderConfig struct {
    // DecodeHook, if set, will be called before any decoding and any
    // type conversion (if WeaklyTypedInput is on). This lets you modify
    // the values before they're set down onto the resulting struct.
    //
    // If an error is returned, the entire decode will fail with that
    // error.
    DecodeHook mapstructure.DecodeHookFunc

    // If ErrorUnused is true, then it is an error for there to exist
    // keys in the original map that were unused in the decoding process
    // (extra keys).
    ErrorUnused bool

    // If WeaklyTypedInput is true, the decoder will make the following
    // "weak" conversions:
    //
    //   - bools to string (true = "1", false = "0")
    //   - numbers to string (base 10)
    //   - bools to int/uint (true = 1, false = 0)
    //   - strings to int/uint (base implied by prefix)
    //   - int to bool (true if value != 0)
    //   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
    //     FALSE, false, False. Anything else is an error)
    //   - empty array = empty map and vice versa
    //
    WeaklyTypedInput bool

    // Metadata is the struct that will contain extra metadata about
    // the decoding. If this is nil, then no metadata will be tracked.
    Metadata *Metadata

    // Result is a pointer to the struct that will contain the decoded
    // value.
    Result interface{}

    // The tag name that mapstructure reads for field names. This
    // defaults to "mapstructure"
    TagName string
}

type Metadata struct {
    // Keys are the keys of the structure which were successfully decoded
    Keys []string

    // Unused is a slice of keys that were found in the raw value but
    // weren't decoded since there was no matching field in the result interface
    Unused []string
}