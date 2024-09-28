package qtypeshttp

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/piotrkowalczuk/qtypes"
	knowntimestamp "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	arraySeparator = ","
	// Null ...
	Null = "null"
	// NotNull ...
	NotNull = "nnull"
	// Equal ...
	Equal = "eq"
	// NotEqual ...
	NotEqual = "neq"
	// GreaterThan ...
	GreaterThan = "gt"
	// NotGreaterThan ...
	NotGreaterThan = "ngt"
	// GreaterThanOrEqual ...
	GreaterThanOrEqual = "gte"
	// NotGreaterThanOrEqual ...
	NotGreaterThanOrEqual = "ngte"
	// LessThan ...
	LessThan = "lt"
	// NotLessThan ...
	NotLessThan = "nlt"
	// LessThanOrEqual ...
	LessThanOrEqual = "lte"
	// NotLessThanOrEqual ...
	NotLessThanOrEqual = "nlte"
	// Between ...
	Between = "bw"
	// NotBetween ...
	NotBetween = "nbw"
	// HasPrefix ...
	HasPrefix = "hp"
	// HasPrefixInsensitive ...
	HasPrefixInsensitive = "hpi"
	// HasSuffix ...
	HasSuffix = "hs"
	// HasSuffixInsensitive ...
	HasSuffixInsensitive = "hsi"
	// Substring ...
	Substring = "sub"
	// SubstringInsensitive` ...
	SubstringInsensitive = "subi"
	// HasElement ...
	HasElement = "he"
	// HasAnyElement ...
	HasAnyElement = "hae"
	// HasAllElements ...
	HasAllElements = "hle"
	// In ...
	In = "in"
	// NotIn ...
	NotIn = "nin"
	// Pattern ...
	Pattern = "rgx"
	// MinLength ...
	MinLength = "minl"
	// MaxLength ...
	MaxLength = "maxl"
	// Contains ...
	Contains = "cts"
	// IsContainedBy ...
	IsContainedBy = "icb"
	// Overlap ...
	Overlap = "ovl"
)

var (
	prefixes = map[string]string{
		Null:                  Null + ":",
		NotNull:               NotNull + ":",
		Equal:                 Equal + ":",
		NotEqual:              NotEqual + ":",
		GreaterThan:           GreaterThan + ":",
		NotGreaterThan:        NotGreaterThan + ":",
		GreaterThanOrEqual:    GreaterThanOrEqual + ":",
		NotGreaterThanOrEqual: NotGreaterThanOrEqual + ":",
		LessThan:              LessThan + ":",
		NotLessThan:           NotLessThan + ":",
		LessThanOrEqual:       LessThanOrEqual + ":",
		NotLessThanOrEqual:    NotLessThanOrEqual + ":",
		Between:               Between + ":",
		NotBetween:            NotBetween + ":",
		HasPrefix:             HasPrefix + ":",
		HasPrefixInsensitive:  HasPrefixInsensitive + ":",
		HasSuffix:             HasSuffix + ":",
		HasSuffixInsensitive:  HasSuffixInsensitive + ":",
		In:                    In + ":",
		NotIn:                 NotIn + ":",
		Substring:             Substring + ":",
		SubstringInsensitive:  SubstringInsensitive + ":",
		Pattern:               Pattern + ":",
		MinLength:             MinLength + ":",
		MaxLength:             MaxLength + ":",
		Contains:              Contains + ":",
		IsContainedBy:         IsContainedBy + ":",
		Overlap:               Overlap + ":",
		HasElement:            HasElement + ":",
		HasAnyElement:         HasAnyElement + ":",
		HasAllElements:        HasAllElements + ":",
	}
)

// ParseInt64 ...
func ParseInt64(s string) (*qtypes.Int64, error) {
	if s == "" {
		return &qtypes.Int64{}, nil
	}
	incoming, t, n, _ := handleNumericPrefix(s)
	outgoing := make([]int64, 0, len(incoming))
	for i, v := range incoming {
		if v == "" {
			break
		}
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("qtypes: query int64 parsing error for key %d and value %v: %s", i, vv, err.Error())
		}
		outgoing = append(outgoing, vv)
	}
	return &qtypes.Int64{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

// ParseFloat64 ...
func ParseFloat64(s string) (*qtypes.Float64, error) {
	if s == "" {
		return &qtypes.Float64{}, nil
	}
	incoming, t, n, _ := handleNumericPrefix(s)

	outgoing := make([]float64, 0, len(incoming))
	for i, v := range incoming {
		if v == "" {
			break
		}
		vv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("qtypes: query float64 parsing error for valur %d: %s", i, err.Error())
		}
		outgoing = append(outgoing, vv)
	}
	return &qtypes.Float64{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

// ParseString allocates new String object based on given string.
// If string is prefixed with known operator e.g. 'hp:New'
// returned object will get same type.
func ParseString(s string) *qtypes.String {
	if s == "" {
		return &qtypes.String{}
	}

	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			t, n, i := queryType(c)
			return &qtypes.String{
				Values:      strings.Split(strings.TrimPrefix(s, p), arraySeparator),
				Type:        t,
				Negation:    n,
				Insensitive: i,
				Valid:       true,
			}
		}
	}
	return &qtypes.String{
		Values: strings.Split(s, arraySeparator),
		Type:   qtypes.QueryType_EQUAL,
		Valid:  true,
	}
}

// ParseTimestamp ...
func ParseTimestamp(s string) (*qtypes.Timestamp, error) {
	if s == "" {
		return &qtypes.Timestamp{}, nil
	}

	incoming, t, n, _ := handleNumericPrefix(s)

	outgoing := make([]*knowntimestamp.Timestamp, 0, len(incoming))
	for i, v := range incoming {
		if v == "" {
			break
		}
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return nil, fmt.Errorf("qtypes: query timestamp parsing error for value %d: %s", i, err.Error())
		}

		outgoing = append(outgoing, knowntimestamp.New(t))
	}
	return &qtypes.Timestamp{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

func queryType(p string) (t qtypes.QueryType, n bool, i bool) {
	switch p {
	case Null:
		t = qtypes.QueryType_NULL
	case NotNull:
		t = qtypes.QueryType_NULL
		n = true
	case Equal:
		t = qtypes.QueryType_EQUAL
	case NotEqual:
		t = qtypes.QueryType_EQUAL
		n = true
	case GreaterThan:
		t = qtypes.QueryType_GREATER
	case NotGreaterThan:
		t = qtypes.QueryType_GREATER
		n = true
	case GreaterThanOrEqual:
		t = qtypes.QueryType_GREATER_EQUAL
	case NotGreaterThanOrEqual:
		t = qtypes.QueryType_GREATER_EQUAL
		n = true
	case LessThan:
		t = qtypes.QueryType_LESS
	case NotLessThan:
		t = qtypes.QueryType_LESS
		n = true
	case LessThanOrEqual:
		t = qtypes.QueryType_LESS_EQUAL
	case NotLessThanOrEqual:
		t = qtypes.QueryType_LESS_EQUAL
		n = true
	case Between:
		t = qtypes.QueryType_BETWEEN
	case NotBetween:
		t = qtypes.QueryType_BETWEEN
		n = true
	case HasElement:
		t = qtypes.QueryType_HAS_ELEMENT
	case HasAllElements:
		t = qtypes.QueryType_HAS_ALL_ELEMENTS
	case HasAnyElement:
		t = qtypes.QueryType_HAS_ANY_ELEMENT
	case HasPrefix:
		t = qtypes.QueryType_HAS_PREFIX
	case HasPrefixInsensitive:
		t = qtypes.QueryType_HAS_PREFIX
		i = true
	case HasSuffix:
		t = qtypes.QueryType_HAS_SUFFIX
	case HasSuffixInsensitive:
		t = qtypes.QueryType_HAS_SUFFIX
		i = true
	case Substring:
		t = qtypes.QueryType_SUBSTRING
	case SubstringInsensitive:
		t = qtypes.QueryType_SUBSTRING
		i = true
	case Pattern:
		t = qtypes.QueryType_PATTERN
	case MinLength:
		t = qtypes.QueryType_MIN_LENGTH
	case MaxLength:
		t = qtypes.QueryType_MAX_LENGTH
	case In:
		t = qtypes.QueryType_IN
	case NotIn:
		t = qtypes.QueryType_IN
		n = true
	case Contains:
		t = qtypes.QueryType_CONTAINS
	case IsContainedBy:
		t = qtypes.QueryType_IS_CONTAINED_BY
	case Overlap:
		t = qtypes.QueryType_OVERLAP
	}
	return
}

func handleNumericPrefix(s string) (incoming []string, t qtypes.QueryType, n, i bool) {
	if parts := strings.Split(s, ":"); len(parts) == 1 {
		return []string{s}, qtypes.QueryType_EQUAL, false, false
	}
	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			t, n, i = queryType(c)
			incoming = strings.Split(strings.TrimPrefix(s, p), arraySeparator)
		}
	}
	if len(incoming) == 0 {
		incoming = strings.Split(s, arraySeparator)
	}

	return
}
