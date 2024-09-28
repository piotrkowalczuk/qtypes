package qtypes

import (
	"fmt"
	"reflect"
	"testing"

	knowntimestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func Example() {
	query := struct {
		Name  *String
		Age   *Int64
		Money *Float64
	}{
		Name: &String{
			Values: []string{"John"},
			Valid:  true,
			Type:   QueryType_SUBSTRING,
		},
		Age:   GreaterInt64(18),
		Money: EqualFloat64(0.0),
	}

	fmt.Println(query.Name.Value())
	fmt.Println(query.Name.Type)
	fmt.Println(query.Age.Value())
	fmt.Println(query.Age.Type)
	fmt.Println(query.Money.Value())
	fmt.Println(query.Money.Type)

	// Output:
	// John
	// SUBSTRING
	// 18
	// GREATER
	// 0
	// EQUAL
}

func ExampleEqualString() {
	ex := EqualString("text")

	fmt.Println(ex.Valid)
	fmt.Println(ex.Negation)
	fmt.Println(ex.Type)
	fmt.Println(ex.Value())

	// Output:
	// true
	// false
	// EQUAL
	// text
}

func TestEqualString(t *testing.T) {
	values := []string{"a"}
	testString(t, EqualString(values[0]), false, true, QueryType_EQUAL, values...)
}

func TestSubString(t *testing.T) {
	values := []string{"a"}
	testString(t, SubString(values[0]), false, true, QueryType_SUBSTRING, values...)
}

func TestHasPrefixString(t *testing.T) {
	values := []string{"a"}
	testString(t, HasPrefixString(values[0]), false, true, QueryType_HAS_PREFIX, values...)
}

func TestHasSuffixString(t *testing.T) {
	values := []string{"a"}
	testString(t, HasSuffixString(values[0]), false, true, QueryType_HAS_SUFFIX, values...)
}

func TestNullString(t *testing.T) {
	testString(t, NullString(), false, true, QueryType_NULL)
}

func testString(t *testing.T, s *String, n, v bool, tp QueryType, values ...string) {
	if s.Negation != n {
		t.Errorf("wrong negation, exiected %t but got %t", n, s.Negation)
	}

	if len(values) > 0 {
		if s.Value() != values[0] {
			t.Errorf("wrong first value, expected %s but got %s", values[0], s.Value())
		}
		if len(s.Values) == len(values) {
			for j, v := range values {
				if s.Values[j] != v {
					t.Errorf("%d: wrong value, expected %s but got %s", j, v, s.Values[j])
				}
			}
		} else {
			t.Errorf("wrong number of values, expected %d but got %d", len(values), len(s.Values))
		}
	}
	if s.Valid != v {
		t.Errorf("expected valid to be %t", v)
	}
	if s.Type != tp {
		t.Errorf("wrong type, expected %s but got %s", tp, s.Type)
	}
}

func TestBetweenTimestamp(t *testing.T) {
	cases := map[string]struct {
		from     *knowntimestamp.Timestamp
		to       *knowntimestamp.Timestamp
		expected Timestamp
	}{
		"valid": {
			from: &knowntimestamp.Timestamp{Seconds: 0, Nanos: 0},
			to:   &knowntimestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: Timestamp{
				Valid:    true,
				Negation: false,
				Type:     QueryType_BETWEEN,
				Values: []*knowntimestamp.Timestamp{
					{Seconds: 0, Nanos: 0},
					{Seconds: 0, Nanos: 1},
				},
			},
		},
		"after-first": {
			from: &knowntimestamp.Timestamp{Seconds: 1, Nanos: 0},
			to:   &knowntimestamp.Timestamp{Seconds: 0, Nanos: 0},
			expected: Timestamp{
				Valid: false,
				Type:  QueryType_BETWEEN,
				Values: []*knowntimestamp.Timestamp{
					{Seconds: 1, Nanos: 0},
					{Seconds: 0, Nanos: 0},
				},
			},
		},
		"after-first-seconds": {
			from: &knowntimestamp.Timestamp{Seconds: 1, Nanos: 1},
			to:   &knowntimestamp.Timestamp{Seconds: 1, Nanos: 0},
			expected: Timestamp{
				Valid: false,
				Type:  QueryType_BETWEEN,
				Values: []*knowntimestamp.Timestamp{
					{Seconds: 1, Nanos: 1},
					{Seconds: 1, Nanos: 0},
				},
			},
		},
		"nil-arguments": {
			from:     nil,
			to:       nil,
			expected: Timestamp{},
		},
		"nil-argument-first": {
			from:     nil,
			to:       &knowntimestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: Timestamp{},
		},
		"nil-argument-second": {
			from:     &knowntimestamp.Timestamp{Seconds: 0, Nanos: 1},
			to:       nil,
			expected: Timestamp{},
		},
	}

	for hint, c := range cases {
		bt := BetweenTimestamp(c.from, c.to)
		if !reflect.DeepEqual(c.expected, *bt) {
			t.Errorf("%s: unexpected output, expected:\n%v\ngot:\n%v\n", hint, c.expected, *bt)
		}
	}
}

func TestInt64_Value(t *testing.T) {
	cases := map[string]struct {
		given    Int64
		expected int64
	}{
		"single": {
			given: Int64{
				Values: []int64{1},
				Valid:  true,
				Type:   QueryType_EQUAL,
			},
			expected: 1,
		},
		"none": {
			given: Int64{
				Valid: true,
				Type:  QueryType_EQUAL,
			},
			expected: 0,
		},
		"multiple": {
			given: Int64{
				Values: []int64{3, 2, 1},
				Valid:  true,
				Type:   QueryType_EQUAL,
			},
			expected: 3,
		},
	}

	for hint, c := range cases {
		if c.given.Value() != c.expected {
			t.Errorf("%s: unexpected value, expected %d but got %d", hint, c.expected, c.given.Value())
		}
	}
}

func TestNaNInt64(t *testing.T) {
	testInt64(t, NullInt64(), false, true, QueryType_NULL)
}

func TestEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, EqualInt64(value), false, true, QueryType_EQUAL, value)
}

func TestNotEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, NotEqualInt64(value), true, true, QueryType_EQUAL, value)
}

func TestGreaterInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, GreaterInt64(value), false, true, QueryType_GREATER, value)
}

func TestGreaterEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, GreaterEqualInt64(value), false, true, QueryType_GREATER_EQUAL, value)
}

func TestBetweenInt64(t *testing.T) {
	values := []int64{1111, 2222}
	testInt64(t, BetweenInt64(values[0], values[1]), false, true, QueryType_BETWEEN, values...)
}

func TestLessInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, LessInt64(value), false, true, QueryType_LESS, value)
}

func TestLessEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, LessEqualInt64(value), false, true, QueryType_LESS_EQUAL, value)
}

func TestInInt64(t *testing.T) {
	values := []int64{1111, 2222, 3333, 4444}
	testInt64(t, InInt64(values...), false, true, QueryType_IN, values...)
}

func testInt64(t *testing.T, i *Int64, n, v bool, tp QueryType, values ...int64) {
	if i.Negation != n {
		t.Errorf("wrong negation, exiected %t but got %t", n, i.Negation)
	}

	if len(values) > 0 {
		if i.Value() != values[0] {
			t.Errorf("wrong first value, expected %d but got %d", values[0], i.Value())
		}
		if len(i.Values) == len(values) {
			for j, v := range values {
				if i.Values[j] != v {
					t.Errorf("%d: wrong value, expected %d but got %d", j, v, i.Values[j])
				}
			}
		} else {
			t.Errorf("wrong number of values, expected %d but got %d", len(values), len(i.Values))
		}
	}
	if i.Valid != v {
		t.Errorf("expected valid to be %t", v)
	}
	if i.Type != tp {
		t.Errorf("wrong type, expected %s but got %s", tp, i.Type)
	}
}

func TestBetweenFloat64(t *testing.T) {
	values := []float64{1111, 2222}
	testFloat64(t, BetweenFloat64(values[0], values[1]), false, true, QueryType_BETWEEN, values...)
}

func testFloat64(t *testing.T, f *Float64, n, v bool, tp QueryType, values ...float64) {
	if f.Negation != n {
		t.Errorf("wrong negation, exiected %t but got %t", n, f.Negation)
	}

	if len(values) > 0 {
		if f.Value() != values[0] {
			t.Errorf("wrong first value, expected %g but got %g", values[0], f.Value())
		}
		if len(f.Values) == len(values) {
			for j, v := range values {
				if f.Values[j] != v {
					t.Errorf("%d: wrong value, expected %g but got %g", j, v, f.Values[j])
				}
			}
		} else {
			t.Errorf("wrong number of values, expected %d but got %d", len(values), len(f.Values))
		}
	}
	if f.Valid != v {
		t.Errorf("expected valid to be %t", v)
	}
	if f.Type != tp {
		t.Errorf("wrong type, expected %s but got %s", tp, f.Type)
	}
}
