// Package provides set of types that helps to build complex protobuf messages that can express conditional statements.
package qtypes

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	pbts "github.com/golang/protobuf/ptypes/timestamp"
)

const (
	arraySeparator = ","
	// Exists ...
	Exists = "ex"
	// ExistsAny ...
	ExistsAny = "exany"
	// ExistsAll ...
	ExistsAll = "exall"
	// NotExists ...
	NotExists = "nex"
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
	// HasSuffix ...
	HasSuffix = "hs"
	// In ...
	In = "in"
	// NotIn ...
	NotIn = "nin"
	// Substring ...
	Substring = "sub"
	// Pattern ...
	Pattern = "rgx"
	// MinLength ...
	MinLength = "minl"
	// MaxLength ...
	MaxLength = "maxl"
	// Any
	Anny = "any"
	// All ...
	All = "all"
	// Contains ...
	Contains = "cts"
	// IsContainedBy ...
	IsContainedBy = "icb"
	// Overlap ...
	Overlap = "ovl"
)

var (
	prefixes = map[string]string{
		Exists:             Exists + ":",
		NotExists:          NotExists + ":",
		Equal:              Equal + ":",
		NotEqual:           NotEqual + ":",
		GreaterThan:        GreaterThan + ":",
		GreaterThanOrEqual: GreaterThanOrEqual + ":",
		LessThan:           LessThan + ":",
		LessThanOrEqual:    LessThanOrEqual + ":",
		Between:            Between + ":",
		NotBetween:         NotBetween + ":",
		HasPrefix:          HasPrefix + ":",
		HasSuffix:          HasSuffix + ":",
		In:                 In + ":",
		Substring:          Substring + ":",
		Pattern:            Pattern + ":",
		MinLength:          MinLength + ":",
		MaxLength:          MaxLength + ":",
	}
)

// Value returns first value or empty string if none.
func (qs *String) Value() string {
	if len(qs.Values) == 0 {
		return ""
	}

	return qs.Values[0]
}

// ParseString allocates new String object based on given string.
// If string is prefixed with known operator e.g. 'hp:New'
// returned object will get same type.
func ParseString(s string) *String {
	if s == "" {
		return &String{}
	}

	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			var (
				t TextQueryType
				n bool
			)
			switch c {
			case Exists:
				t = TextQueryType_NOT_A_TEXT
				n = true
			case NotExists:
				t = TextQueryType_NOT_A_TEXT
			case Equal:
				t = TextQueryType_EXACT
			case NotEqual:
				t = TextQueryType_EXACT
				n = true
			case HasPrefix:
				t = TextQueryType_HAS_PREFIX
			case HasSuffix:
				t = TextQueryType_HAS_SUFFIX
			case Substring:
				t = TextQueryType_SUBSTRING
			case Pattern:
				t = TextQueryType_PATTERN
			case MinLength:
				t = TextQueryType_MIN_LENGTH
			case MaxLength:
				t = TextQueryType_MAX_LENGTH
			}
			return &String{
				Values:   strings.Split(strings.TrimLeft(s, p), arraySeparator),
				Type:     t,
				Negation: n,
				Valid:    true,
			}
		}
	}
	return &String{
		Values: strings.Split(s, arraySeparator),
		Type:   TextQueryType_EXACT,
		Valid:  true,
	}
}

// ExactString ...
func ExactString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_EXACT,
	}
}

// HasPrefixString ...
func HasPrefixString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_HAS_PREFIX,
	}
}

// HasSuffixString ...
func HasSuffixString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_HAS_SUFFIX,
	}
}

// SubString ...
func SubString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_SUBSTRING,
	}
}

// NotATextString ...
func NotATextString() *String {
	return &String{
		Valid: true,
		Type:  TextQueryType_NOT_A_TEXT,
	}
}

// NaNInt64 allocates valid Int64 object of type not a number with given value.
func NaNInt64() *Int64 {
	return &Int64{
		Valid:  true,
		Type:   NumericQueryType_NOT_A_NUMBER,
	}
}

// EqualInt64 allocates valid Int64 object of type equal with given value.
func EqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// NotEqualInt64 allocates valid Int64 negated object of type equal with given value.
func NotEqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Negation:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// InInt64 allocates valid Int64 object of type in with given values.
func InInt64(v ...int64) *Int64 {
	return &Int64{
		Values: v,
		Valid:  true,
		Type:   NumericQueryType_IN,
	}
}

// BetweenInt64 allocates valid Int64 object of type between with given values.
func BetweenInt64(a, b int64) *Int64 {
	return &Int64{
		Values: []int64{a, b},
		Valid:  true,
		Type:   NumericQueryType_BETWEEN,
	}
}

// GreaterInt64 allocates valid Int64 object of type greater with given value.
func GreaterInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_GREATER,
	}
}

// GreaterEqualInt64 allocates valid Int64 object of type greater equal with given value.
func GreaterEqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_GREATER_EQUAL,
	}
}

// LessInt64 allocates valid Int64 object of type less with given value.
func LessInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_LESS,
	}
}

// LessEqualInt64 allocates valid Int64 object of type less equal with given value.
func LessEqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_LESS_EQUAL,
	}
}

// Value ...
func (qi *Int64) Value() int64 {
	if len(qi.Values) == 0 {
		return 0
	}

	return qi.Values[0]
}

func ParseInt64(s string) (*Int64, error) {
	if s == "" {
		return &Int64{}, nil
	}
	var (
		t        NumericQueryType
		n        bool
		incoming []string
	)
	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			switch c {
			case Exists:
				t = NumericQueryType_NOT_A_NUMBER
				n = true
			case NotExists:
				t = NumericQueryType_NOT_A_NUMBER
			case Equal:
				t = NumericQueryType_EQUAL
			case NotEqual:
				t = NumericQueryType_EQUAL
				n = true
			case GreaterThan:
				t = NumericQueryType_GREATER
			case GreaterThanOrEqual:
				t = NumericQueryType_GREATER_EQUAL
			case LessThan:
				t = NumericQueryType_LESS
			case LessThanOrEqual:
				t = NumericQueryType_LESS_EQUAL
			case Between:
				t = NumericQueryType_BETWEEN
			case NotBetween:
				t = NumericQueryType_BETWEEN
				n = true
			}

			fmt.Println(s, p, strings.TrimLeft(s, p))
			incoming = strings.Split(strings.TrimLeft(s, p), arraySeparator)

		}
	}
	if len(incoming) == 0 {
		incoming = strings.Split(s, arraySeparator)

	}

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
	return &Int64{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

// EqualFloat64 allocates valid Float64 object of type equal with given value.
func EqualFloat64(i float64) *Float64 {
	return &Float64{
		Values: []float64{i},
		Valid:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// BetweenFloat64 allocates valid Float64 object if both numbers are not 0 and from is not greater than to.
func BetweenFloat64(from, to float64) *Float64 {
	if from == 0 && to == 0 {
		return &Float64{}
	}
	if from > to {
		return &Float64{}
	}
	return &Float64{
		Values: []float64{from, to},
		Type:   NumericQueryType_BETWEEN,
		Valid:  true,
	}
}

// Value returns first available value or 0 if none available.
func (qf *Float64) Value() float64 {
	if len(qf.Values) == 0 {
		return 0.0
	}

	return qf.Values[0]
}

// ParseFloat64 ...
func ParseFloat64(s string) (*Float64, error) {
	if s == "" {
		return &Float64{}, nil
	}
	incoming, t, n := handleNumericPrefix(s)

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
	return &Float64{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

// BetweenTimestamp allocates valid Timestamp object if both timestamps are not nil
// and first is before the second.
func BetweenTimestamp(from, to *pbts.Timestamp) *Timestamp {
	if from == nil || to == nil {
		return &Timestamp{}
	}

	v := true
	if to.Seconds < from.Seconds {
		v = false
	}
	if to.Seconds == from.Seconds && to.Nanos < from.Nanos {
		v = false
	}
	return &Timestamp{
		Values: []*pbts.Timestamp{from, to},
		Type:   NumericQueryType_BETWEEN,
		Valid:  v,
	}
}

// Value returns first value or nil if none.
func (qt *Timestamp) Value() *pbts.Timestamp {
	if len(qt.Values) == 0 {
		return nil
	}

	return qt.Values[0]
}

// ParseTimestamp ...
func ParseTimestamp(s string) (*Timestamp, error) {
	if s == "" {
		return &Timestamp{}, nil
	}

	incoming, t, n := handleNumericPrefix(s)

	outgoing := make([]*pbts.Timestamp, 0, len(incoming))
	for i, v := range incoming {
		if v == "" {
			break
		}
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return nil, fmt.Errorf("qtypes: query timestamp parsing error for value %d: %s", i, err.Error())
		}
		tt, err := ptypes.TimestampProto(t)
		if err != nil {
			return nil, fmt.Errorf("qtypes: time to proto timestamp conversion error for value %d: %s", i, err.Error())
		}
		outgoing = append(outgoing, tt)
	}
	return &Timestamp{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

func handleNumericPrefix(s string) (incoming []string, t NumericQueryType, n bool) {
	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			switch c {
			case Exists:
				t = NumericQueryType_NOT_A_NUMBER
				n = true
			case NotExists:
				t = NumericQueryType_NOT_A_NUMBER
			case Equal:
				t = NumericQueryType_EQUAL
			case NotEqual:
				t = NumericQueryType_EQUAL
				n = true
			case GreaterThan:
				t = NumericQueryType_GREATER
			case NotGreaterThan:
				t = NumericQueryType_GREATER
				n = true
			case GreaterThanOrEqual:
				t = NumericQueryType_GREATER_EQUAL
			case NotGreaterThanOrEqual:
				t = NumericQueryType_GREATER_EQUAL
				n = true
			case LessThan:
				t = NumericQueryType_LESS
			case NotLessThan:
				t = NumericQueryType_LESS
				n = true
			case LessThanOrEqual:
				t = NumericQueryType_LESS_EQUAL
			case NotLessThanOrEqual:
				t = NumericQueryType_LESS_EQUAL
				n = true
			case Between:
				t = NumericQueryType_BETWEEN
			case NotBetween:
				t = NumericQueryType_BETWEEN
				n = true
			case In:
				t = NumericQueryType_IN
			case NotIn:
				t = NumericQueryType_IN
				n = true
			}

			incoming = strings.Split(strings.TrimLeft(s, p), arraySeparator)
		}
	}
	if len(incoming) == 0 {
		incoming = strings.Split(s, arraySeparator)
	}

	return
}
