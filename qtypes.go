// Package qtypes provides set of types that helps to build complex protobuf messages that can express conditional statements.
package qtypes

import (
	pbts "github.com/golang/protobuf/ptypes/timestamp"
)

// Value returns first value or empty string if none.
func (qs *String) Value() string {
	if len(qs.Values) == 0 {
		return ""
	}

	return qs.Values[0]
}

// EqualString ...
func EqualString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   QueryType_EQUAL,
	}
}

// HasPrefixString ...
func HasPrefixString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   QueryType_HAS_PREFIX,
	}
}

// HasSuffixString ...
func HasSuffixString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   QueryType_HAS_SUFFIX,
	}
}

// SubString ...
func SubString(s string) *String {
	return &String{
		Values: []string{s},
		Valid:  true,
		Type:   QueryType_SUBSTRING,
	}
}

// NullString ...
func NullString() *String {
	return &String{
		Valid: true,
		Type:  QueryType_NULL,
	}
}

// NullInt64 allocates valid Int64 object of type not a number with given value.
func NullInt64() *Int64 {
	return &Int64{
		Valid: true,
		Type:  QueryType_NULL,
	}
}

// EqualInt64 allocates valid Int64 object of type equal with given value.
func EqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   QueryType_EQUAL,
	}
}

// NotEqualInt64 allocates valid Int64 negated object of type equal with given value.
func NotEqualInt64(i int64) *Int64 {
	return &Int64{
		Values:   []int64{i},
		Valid:    true,
		Negation: true,
		Type:     QueryType_EQUAL,
	}
}

// InInt64 allocates valid Int64 object of type in with given values.
func InInt64(v ...int64) *Int64 {
	return &Int64{
		Values: v,
		Valid:  true,
		Type:   QueryType_IN,
	}
}

// BetweenInt64 allocates valid Int64 object of type between with given values.
func BetweenInt64(a, b int64) *Int64 {
	return &Int64{
		Values: []int64{a, b},
		Valid:  true,
		Type:   QueryType_BETWEEN,
	}
}

// GreaterInt64 allocates valid Int64 object of type greater with given value.
func GreaterInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   QueryType_GREATER,
	}
}

// GreaterEqualInt64 allocates valid Int64 object of type greater equal with given value.
func GreaterEqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   QueryType_GREATER_EQUAL,
	}
}

// LessInt64 allocates valid Int64 object of type less with given value.
func LessInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   QueryType_LESS,
	}
}

// LessEqualInt64 allocates valid Int64 object of type less equal with given value.
func LessEqualInt64(i int64) *Int64 {
	return &Int64{
		Values: []int64{i},
		Valid:  true,
		Type:   QueryType_LESS_EQUAL,
	}
}

// Value ...
func (i *Int64) Value() int64 {
	if len(i.Values) == 0 {
		return 0
	}

	return i.Values[0]
}

// EqualFloat64 allocates valid Float64 object of type equal with given value.
func EqualFloat64(i float64) *Float64 {
	return &Float64{
		Values: []float64{i},
		Valid:  true,
		Type:   QueryType_EQUAL,
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
		Type:   QueryType_BETWEEN,
		Valid:  true,
	}
}

// Value returns first available value or 0 if none available.
func (f *Float64) Value() float64 {
	if len(f.Values) == 0 {
		return 0.0
	}

	return f.Values[0]
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
		Type:   QueryType_BETWEEN,
		Valid:  v,
	}
}

// Value returns first value or nil if none.
func (t *Timestamp) Value() *pbts.Timestamp {
	if len(t.Values) == 0 {
		return nil
	}

	return t.Values[0]
}
