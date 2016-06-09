package qtypes

import (
	"fmt"
	"reflect"
	"testing"

	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func Example() {
	query := struct {
		Name  *String
		Age   *Int64
		Money *Float64
	}{
		Name:  ParseString("sub:John"),
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

func ExampleExactString() {
	ex := ExactString("text")

	fmt.Println(ex.Valid)
	fmt.Println(ex.Negation)
	fmt.Println(ex.Type)
	fmt.Println(ex.Value())

	// Output:
	// true
	// false
	// EXACT
	// text
}

func TestParseString(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected String
	}{
		"exists": {
			given: "ex:",
			expected: String{
				Values:   []string{""},
				Type:     TextQueryType_NOT_A_TEXT,
				Valid:    true,
				Negation: true,
			},
		},
		"not-exists": {
			given: "nex:",
			expected: String{
				Values: []string{""},
				Type:   TextQueryType_NOT_A_TEXT,
				Valid:  true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: String{
				Values: []string{"123"},
				Type:   TextQueryType_EXACT,
				Valid:  true,
			},
		},
		"has-prefix": {
			given: "hp:New",
			expected: String{
				Values: []string{"New"},
				Type:   TextQueryType_HAS_PREFIX,
				Valid:  true,
			},
		},
		"has-suffix": {
			given: "hs:New",
			expected: String{
				Values: []string{"New"},
				Type:   TextQueryType_HAS_SUFFIX,
				Valid:  true,
			},
		},
		"substring": {
			given: "sub:anything",
			expected: String{
				Values: []string{"anything"},
				Type:   TextQueryType_SUBSTRING,
				Valid:  true,
			},
		},
		"pattern": {
			given: "rgx:.*",
			expected: String{
				Values: []string{".*"},
				Type:   TextQueryType_PATTERN,
				Valid:  true,
			},
		},
		"max-length": {
			given: "maxl:4",
			expected: String{
				Values: []string{"4"},
				Type:   TextQueryType_MAX_LENGTH,
				Valid:  true,
			},
		},
		"min-length": {
			given: "minl:555",
			expected: String{
				Values: []string{"555"},
				Type:   TextQueryType_MIN_LENGTH,
				Valid:  true,
			},
		},
		"empty": {
			given:    "",
			expected: String{},
		},
		"without-condition": {
			given: "text",
			expected: String{
				Values: []string{"text"},
				Type:   TextQueryType_EXACT,
				Valid:  true,
			},
		},
		"with-condition-but-without-value": {
			given: "neq:",
			expected: String{
				Values:   []string{""},
				Type:     TextQueryType_EXACT,
				Valid:    true,
				Negation: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got := ParseString(c.given)

		if got == nil {
			t.Errorf("unexpected nil")
			continue CasesLoop
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestExactString(t *testing.T) {
	es := ExactString("John")

	if es.Negation {
		t.Errorf("unexpected negation")
	}
	if es.Value() != "John" {
		t.Errorf("unexpected value")
	}
	if !es.Valid {
		t.Errorf("expected to be valid")
	}
}

func TestBetweenTimestamp(t *testing.T) {
	cases := map[string]struct {
		from     *timestamp.Timestamp
		to       *timestamp.Timestamp
		expected Timestamp
	}{
		"valid": {
			from: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
			to:   &timestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: Timestamp{
				Valid:    true,
				Negation: false,
				Type:     NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 0, Nanos: 0},
					&timestamp.Timestamp{Seconds: 0, Nanos: 1},
				},
			},
		},
		"after-first": {
			from: &timestamp.Timestamp{Seconds: 1, Nanos: 0},
			to:   &timestamp.Timestamp{Seconds: 0, Nanos: 0},
			expected: Timestamp{
				Valid: false,
				Type:  NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 1, Nanos: 0},
					&timestamp.Timestamp{Seconds: 0, Nanos: 0},
				},
			},
		},
		"after-first-seconds": {
			from: &timestamp.Timestamp{Seconds: 1, Nanos: 1},
			to:   &timestamp.Timestamp{Seconds: 1, Nanos: 0},
			expected: Timestamp{
				Valid: false,
				Type:  NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 1, Nanos: 1},
					&timestamp.Timestamp{Seconds: 1, Nanos: 0},
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
			to:       &timestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: Timestamp{},
		},
		"nil-argument-second": {
			from:     &timestamp.Timestamp{Seconds: 0, Nanos: 1},
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
				Type:   NumericQueryType_EQUAL,
			},
			expected: 1,
		},
		"none": {
			given: Int64{
				Valid: true,
				Type:  NumericQueryType_EQUAL,
			},
			expected: 0,
		},
		"multiple": {
			given: Int64{
				Values: []int64{3, 2, 1},
				Valid:  true,
				Type:   NumericQueryType_EQUAL,
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
	testInt64(t, NaNInt64(), false, true, NumericQueryType_NOT_A_NUMBER)
}

func TestEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, EqualInt64(value), false, true, NumericQueryType_EQUAL, value)
}

func TestNotEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, NotEqualInt64(value), true, true, NumericQueryType_EQUAL, value)
}

func TestGreaterInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, GreaterInt64(value), false, true, NumericQueryType_GREATER, value)
}

func TestGreaterEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, GreaterEqualInt64(value), false, true, NumericQueryType_GREATER_EQUAL, value)
}

func TestBetweenInt64(t *testing.T) {
	values := []int64{1111, 2222}
	testInt64(t, BetweenInt64(values[0], values[1]), false, true, NumericQueryType_BETWEEN, values...)
}

func TestLessInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, LessInt64(value), false, true, NumericQueryType_LESS, value)
}

func TestLessEqualInt64(t *testing.T) {
	value := int64(1111)
	testInt64(t, LessEqualInt64(value), false, true, NumericQueryType_LESS_EQUAL, value)
}

func TestInInt64(t *testing.T) {
	values := []int64{1111, 2222, 3333, 4444}
	testInt64(t, InInt64(values...), false, true, NumericQueryType_IN, values...)
}

func TestParseInt64(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected Int64
	}{
		"empty": {
			given:    "",
			expected: Int64{},
		},
		"exists": {
			given: "ex:",
			expected: Int64{
				Values:   []int64{},
				Type:     NumericQueryType_NOT_A_NUMBER,
				Valid:    true,
				Negation: true,
			},
		},
		"not-exists": {
			given: "nex:",
			expected: Int64{
				Values: []int64{},
				Type:   NumericQueryType_NOT_A_NUMBER,
				Valid:  true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: Int64{
				Values: []int64{123},
				Type:   NumericQueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-equal": {
			given: "neq:123",
			expected: Int64{
				Values:   []int64{123},
				Type:     NumericQueryType_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"greater": {
			given: "gt:555",
			expected: Int64{
				Values: []int64{555},
				Type:   NumericQueryType_GREATER,
				Valid:  true,
			},
		},
		"greater-equal": {
			given: "gte:666",
			expected: Int64{
				Values: []int64{666},
				Type:   NumericQueryType_GREATER_EQUAL,
				Valid:  true,
			},
		},
		"lesser": {
			given: "lt:777",
			expected: Int64{
				Values: []int64{777},
				Type:   NumericQueryType_LESS,
				Valid:  true,
			},
		},
		"lesser-equal": {
			given: "lte:888",
			expected: Int64{
				Values: []int64{888},
				Type:   NumericQueryType_LESS_EQUAL,
				Valid:  true,
			},
		},
		"between": {
			given: "bw:111,222",
			expected: Int64{
				Values: []int64{111, 222},
				Type:   NumericQueryType_BETWEEN,
				Valid:  true,
			},
		},
		"not-between": {
			given: "nbw:111,222",
			expected: Int64{
				Values:   []int64{111, 222},
				Type:     NumericQueryType_BETWEEN,
				Valid:    true,
				Negation: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got, err := ParseInt64(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", hint, err.Error())
			continue CasesLoop
		}
		if got == nil {
			t.Errorf("%s: unexpected nil", hint)
			continue CasesLoop
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestParseInt64_text(t *testing.T) {
	got, err := ParseInt64("ne:long-text")
	if err == nil {
		t.Fatalf("expected error")
	}
	if got != nil {
		t.Fatalf("expected nil")
	}
}
func TestParseTimestamp(t *testing.T) {
	parseTimestamp := func(t *testing.T, s string) *timestamp.Timestamp {
		pt, err := time.Parse(time.RFC3339Nano, s)
		if err != nil {
			t.Fatalf("string cant be parsed into time: %s", err.Error())
		}
		tt, err := ptypes.TimestampProto(pt)
		if err != nil {
			t.Fatalf("tmie cant be converted into timestamp: %s", err.Error())
		}
		return tt
	}
	cases := map[string]struct {
		given    string
		expected Timestamp
	}{
		"empty": {
			given:    "",
			expected: Timestamp{},
		},
		"exists": {
			given: "ex:",
			expected: Timestamp{
				Values:   []*timestamp.Timestamp{},
				Type:     NumericQueryType_NOT_A_NUMBER,
				Valid:    true,
				Negation: true,
			},
		},
		"not-exists": {
			given: "nex:",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{},
				Type:   NumericQueryType_NOT_A_NUMBER,
				Valid:  true,
			},
		},
		"equal": {
			given: "eq:2009-11-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  NumericQueryType_EQUAL,
				Valid: true,
			},
		},
		"greater-equal": {
			given: "gte:2009-11-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  NumericQueryType_GREATER_EQUAL,
				Valid: true,
			},
		},
		"greater": {
			given: "gt:2009-11-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  NumericQueryType_GREATER,
				Valid: true,
			},
		},
		"less": {
			given: "lt:2009-11-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  NumericQueryType_LESS,
				Valid: true,
			},
		},
		"less-equal": {
			given: "lte:2009-11-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  NumericQueryType_LESS_EQUAL,
				Valid: true,
			},
		},
		"between": {
			given: "bw:2009-11-10T23:00:00Z,2009-12-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
					parseTimestamp(t, "2009-12-10T23:00:00Z"),
				},
				Type:  NumericQueryType_BETWEEN,
				Valid: true,
			},
		},
		"in": {
			given: "in:2009-10-10T23:00:00Z,2009-11-10T23:00:00Z,2009-12-10T23:00:00Z",
			expected: Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-10-10T23:00:00Z"),
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
					parseTimestamp(t, "2009-12-10T23:00:00Z"),
				},
				Type:  NumericQueryType_IN,
				Valid: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got, err := ParseTimestamp(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", hint, err.Error())
			continue CasesLoop
		}
		if got == nil {
			t.Errorf("%s: unexpected nil", hint)
			continue CasesLoop
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func testInt64(t *testing.T, i *Int64, n, v bool, tp NumericQueryType, values ...int64) {
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
