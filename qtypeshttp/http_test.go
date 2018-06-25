package qtypeshttp_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/piotrkowalczuk/qtypes"
	"github.com/piotrkowalczuk/qtypes/qtypeshttp"
)

func TestParseString(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected qtypes.String
	}{
		"null": {
			given: "null:",
			expected: qtypes.String{
				Values: []string{""},
				Type:   qtypes.QueryType_NULL,
				Valid:  true,
			},
		},
		"not-null": {
			given: "nnull:",
			expected: qtypes.String{
				Values:   []string{""},
				Type:     qtypes.QueryType_NULL,
				Valid:    true,
				Negation: true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: qtypes.String{
				Values: []string{"123"},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"has-prefix": {
			given: "hp:New",
			expected: qtypes.String{
				Values: []string{"New"},
				Type:   qtypes.QueryType_HAS_PREFIX,
				Valid:  true,
			},
		},
		"has-prefix-insensitive": {
			given: "hpi:New",
			expected: qtypes.String{
				Values:      []string{"New"},
				Type:        qtypes.QueryType_HAS_PREFIX,
				Valid:       true,
				Insensitive: true,
			},
		},
		"has-suffix": {
			given: "hs:New",
			expected: qtypes.String{
				Values: []string{"New"},
				Type:   qtypes.QueryType_HAS_SUFFIX,
				Valid:  true,
			},
		},
		"has-suffix-insensitive": {
			given: "hsi:New",
			expected: qtypes.String{
				Values:      []string{"New"},
				Type:        qtypes.QueryType_HAS_SUFFIX,
				Valid:       true,
				Insensitive: true,
			},
		},
		"substring": {
			given: "sub:anything",
			expected: qtypes.String{
				Values: []string{"anything"},
				Type:   qtypes.QueryType_SUBSTRING,
				Valid:  true,
			},
		},
		"substring-insensitive": {
			given: "subi:anything",
			expected: qtypes.String{
				Values:      []string{"anything"},
				Type:        qtypes.QueryType_SUBSTRING,
				Valid:       true,
				Insensitive: true,
			},
		},
		"pattern": {
			given: "rgx:.*",
			expected: qtypes.String{
				Values: []string{".*"},
				Type:   qtypes.QueryType_PATTERN,
				Valid:  true,
			},
		},
		"max-length": {
			given: "maxl:4",
			expected: qtypes.String{
				Values: []string{"4"},
				Type:   qtypes.QueryType_MAX_LENGTH,
				Valid:  true,
			},
		},
		"min-length": {
			given: "minl:555",
			expected: qtypes.String{
				Values: []string{"555"},
				Type:   qtypes.QueryType_MIN_LENGTH,
				Valid:  true,
			},
		},
		"empty": {
			given:    "",
			expected: qtypes.String{},
		},
		"without-condition": {
			given: "text",
			expected: qtypes.String{
				Values: []string{"text"},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"with-condition-but-without-value": {
			given: "neq:",
			expected: qtypes.String{
				Values:   []string{""},
				Type:     qtypes.QueryType_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"has-element": {
			given: "he:555",
			expected: qtypes.String{
				Values: []string{"555"},
				Type:   qtypes.QueryType_HAS_ELEMENT,
				Valid:  true,
			},
		},
		"has-any-elements": {
			given: "hae:555,222",
			expected: qtypes.String{
				Values: []string{"555", "222"},
				Type:   qtypes.QueryType_HAS_ANY_ELEMENT,
				Valid:  true,
			},
		},
		"has-all-elements": {
			given: "hle:111,222",
			expected: qtypes.String{
				Values: []string{"111", "222"},
				Type:   qtypes.QueryType_HAS_ALL_ELEMENTS,
				Valid:  true,
			},
		},
		"not-greater-than": {
			given: "ngt:111",
			expected: qtypes.String{
				Values:   []string{"111"},
				Type:     qtypes.QueryType_GREATER,
				Negation: true,
				Valid:    true,
			},
		},
		"greater-than": {
			given: "gt:111",
			expected: qtypes.String{
				Values: []string{"111"},
				Type:   qtypes.QueryType_GREATER,
				Valid:  true,
			},
		},
		"less-than": {
			given: "lt:111",
			expected: qtypes.String{
				Values: []string{"111"},
				Type:   qtypes.QueryType_LESS,
				Valid:  true,
			},
		},
	}

	for hint, c := range cases {
		got := qtypeshttp.ParseString(c.given)

		if got == nil {
			t.Error("unexpected nil")
			continue
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestParseFloat64(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected qtypes.Float64
	}{
		"empty": {
			given:    "",
			expected: qtypes.Float64{},
		},
		"null": {
			given: "null:",
			expected: qtypes.Float64{
				Values: []float64{},
				Type:   qtypes.QueryType_NULL,
				Valid:  true,
			},
		},
		"number": {
			given: "15.15",
			expected: qtypes.Float64{
				Values: []float64{15.15},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-null": {
			given: "nnull:",
			expected: qtypes.Float64{
				Values:   []float64{},
				Type:     qtypes.QueryType_NULL,
				Valid:    true,
				Negation: true,
			},
		},
		"equal": {
			given: "eq:123.15",
			expected: qtypes.Float64{
				Values: []float64{123.15},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-equal": {
			given: "neq:123.55555",
			expected: qtypes.Float64{
				Values:   []float64{123.55555},
				Type:     qtypes.QueryType_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"greater": {
			given: "gt:555",
			expected: qtypes.Float64{
				Values: []float64{555.00},
				Type:   qtypes.QueryType_GREATER,
				Valid:  true,
			},
		},
		"greater-equal": {
			given: "gte:666.666",
			expected: qtypes.Float64{
				Values: []float64{666.666},
				Type:   qtypes.QueryType_GREATER_EQUAL,
				Valid:  true,
			},
		},
		"lesser": {
			given: "lt:777.666",
			expected: qtypes.Float64{
				Values: []float64{777.666},
				Type:   qtypes.QueryType_LESS,
				Valid:  true,
			},
		},
		"lesser-equal": {
			given: "lte:888.666",
			expected: qtypes.Float64{
				Values: []float64{888.666},
				Type:   qtypes.QueryType_LESS_EQUAL,
				Valid:  true,
			},
		},
		"between": {
			given: "bw:111.666,222.666",
			expected: qtypes.Float64{
				Values: []float64{111.666, 222.666},
				Type:   qtypes.QueryType_BETWEEN,
				Valid:  true,
			},
		},
		"not-between": {
			given: "nbw:111.666,222",
			expected: qtypes.Float64{
				Values:   []float64{111.666, 222},
				Type:     qtypes.QueryType_BETWEEN,
				Valid:    true,
				Negation: true,
			},
		},
		"not-less": {
			given: "nlt:111.666",
			expected: qtypes.Float64{
				Values:   []float64{111.666},
				Type:     qtypes.QueryType_LESS,
				Valid:    true,
				Negation: true,
			},
		},
		"not-greater-than-or-equal": {
			given: "ngte:111.666",
			expected: qtypes.Float64{
				Values:   []float64{111.666},
				Type:     qtypes.QueryType_GREATER_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"not-less-than-or-equal": {
			given: "nlte:111.666",
			expected: qtypes.Float64{
				Values:   []float64{111.666},
				Type:     qtypes.QueryType_LESS_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"not-in": {
			given: "nin:111.666,222.444",
			expected: qtypes.Float64{
				Values:   []float64{111.666, 222.444},
				Type:     qtypes.QueryType_IN,
				Valid:    true,
				Negation: true,
			},
		},
		"contains": {
			given: "cts:111.666,222.444",
			expected: qtypes.Float64{
				Values: []float64{111.666, 222.444},
				Type:   qtypes.QueryType_CONTAINS,
				Valid:  true,
			},
		},
		"is-contained-by": {
			given: "icb:111.666,222.444",
			expected: qtypes.Float64{
				Values: []float64{111.666, 222.444},
				Type:   qtypes.QueryType_IS_CONTAINED_BY,
				Valid:  true,
			},
		},
		"overlap": {
			given: "ovl:111.666,222.444",
			expected: qtypes.Float64{
				Values: []float64{111.666, 222.444},
				Type:   qtypes.QueryType_OVERLAP,
				Valid:  true,
			},
		},
	}

	for hint, c := range cases {
		got, err := qtypeshttp.ParseFloat64(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", hint, err.Error())
			continue
		}
		if got == nil {
			t.Errorf("%s: unexpected nil", hint)
			continue
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
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
		expected qtypes.Timestamp
	}{
		"empty": {
			given:    "",
			expected: qtypes.Timestamp{},
		},
		"null": {
			given: "null:",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{},
				Type:   qtypes.QueryType_NULL,
				Valid:  true,
			},
		},
		"not-null": {
			given: "nnull:",
			expected: qtypes.Timestamp{
				Values:   []*timestamp.Timestamp{},
				Type:     qtypes.QueryType_NULL,
				Valid:    true,
				Negation: true,
			},
		},
		"equal": {
			given: "eq:2009-11-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_EQUAL,
				Valid: true,
			},
		},
		"greater-equal": {
			given: "gte:2009-11-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_GREATER_EQUAL,
				Valid: true,
			},
		},
		"greater": {
			given: "gt:2009-11-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_GREATER,
				Valid: true,
			},
		},
		"less": {
			given: "lt:2009-11-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_LESS,
				Valid: true,
			},
		},
		"less-equal": {
			given: "lte:2009-11-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_LESS_EQUAL,
				Valid: true,
			},
		},
		"between": {
			given: "bw:2009-11-10T23:00:00Z,2009-12-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
					parseTimestamp(t, "2009-12-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_BETWEEN,
				Valid: true,
			},
		},
		"in": {
			given: "in:2009-10-10T23:00:00Z,2009-11-10T23:00:00Z,2009-12-10T23:00:00Z",
			expected: qtypes.Timestamp{
				Values: []*timestamp.Timestamp{
					parseTimestamp(t, "2009-10-10T23:00:00Z"),
					parseTimestamp(t, "2009-11-10T23:00:00Z"),
					parseTimestamp(t, "2009-12-10T23:00:00Z"),
				},
				Type:  qtypes.QueryType_IN,
				Valid: true,
			},
		},
	}

	for hint, c := range cases {
		got, err := qtypeshttp.ParseTimestamp(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", hint, err.Error())
			continue
		}
		if got == nil {
			t.Errorf("%s: unexpected nil", hint)
			continue
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestParseInt64(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected qtypes.Int64
	}{
		"empty": {
			given:    "",
			expected: qtypes.Int64{},
		},
		"null": {
			given: "null:",
			expected: qtypes.Int64{
				Values: []int64{},
				Type:   qtypes.QueryType_NULL,
				Valid:  true,
			},
		},
		"number": {
			given: "15",
			expected: qtypes.Int64{
				Values: []int64{15},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-null": {
			given: "nnull:",
			expected: qtypes.Int64{
				Values:   []int64{},
				Type:     qtypes.QueryType_NULL,
				Valid:    true,
				Negation: true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: qtypes.Int64{
				Values: []int64{123},
				Type:   qtypes.QueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-equal": {
			given: "neq:123",
			expected: qtypes.Int64{
				Values:   []int64{123},
				Type:     qtypes.QueryType_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"greater": {
			given: "gt:555",
			expected: qtypes.Int64{
				Values: []int64{555},
				Type:   qtypes.QueryType_GREATER,
				Valid:  true,
			},
		},
		"greater-equal": {
			given: "gte:666",
			expected: qtypes.Int64{
				Values: []int64{666},
				Type:   qtypes.QueryType_GREATER_EQUAL,
				Valid:  true,
			},
		},
		"lesser": {
			given: "lt:777",
			expected: qtypes.Int64{
				Values: []int64{777},
				Type:   qtypes.QueryType_LESS,
				Valid:  true,
			},
		},
		"lesser-equal": {
			given: "lte:888",
			expected: qtypes.Int64{
				Values: []int64{888},
				Type:   qtypes.QueryType_LESS_EQUAL,
				Valid:  true,
			},
		},
		"between": {
			given: "bw:111,222",
			expected: qtypes.Int64{
				Values: []int64{111, 222},
				Type:   qtypes.QueryType_BETWEEN,
				Valid:  true,
			},
		},
		"not-between": {
			given: "nbw:111,222",
			expected: qtypes.Int64{
				Values:   []int64{111, 222},
				Type:     qtypes.QueryType_BETWEEN,
				Valid:    true,
				Negation: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got, err := qtypeshttp.ParseInt64(c.given)
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
	got, err := qtypeshttp.ParseInt64("ne:long-text")
	if err == nil {
		t.Fatal("expected error")
	}
	if got != nil {
		t.Fatal("expected nil")
	}
}
