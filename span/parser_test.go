package span

import (
	"testing"
)

func TestParseTimestamp(t *testing.T) {

	tests := map[string]struct {
		expected   Timestamp
		input, err string
	}{
		`empty`: {
			expected: Timestamp{},
			input:    "",
			err:      "timestamp must be in the format hh:mm",
		},
		`simple`: {
			expected: Timestamp{10, 0},
			input:    "10:00",
			err:      "",
		},
		`spaced`: {
			expected: Timestamp{9, 15},
			input:    " 9:15",
			err:      "",
		},
		`halfa`: {
			expected: Timestamp{14, 30},
			input:    "14:30",
			err:      "",
		},
		`hightide`: {
			expected: Timestamp{},
			input:    "24:30",
			err:      "hours must be between 0 and 23",
		},
		`negative`: {
			expected: Timestamp{},
			input:    "-9:15",
			err:      "hours must be between 0 and 23",
		},
		`negativeMinutes`: {
			expected: Timestamp{},
			input:    "12:-5",
			err:      "minutes must be 0, 15, 30 or 45",
		},
		`wrong`: {
			expected: Timestamp{},
			input:    "not a timestamp",
			err:      "timestamp must be in the format hh:mm",
		},
		`malf`: {
			expected: Timestamp{},
			input:    "3X:12",
			err:      "unable to parse hours",
		},
		`malfMinutes`: {
			expected: Timestamp{},
			input:    "17:I2",
			err:      "unable to parse minutes",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseTimestamp(tt.input)
			if err != nil {
				if err.Error() != tt.err {
					t.Errorf("got: %s, expected: %s", err.Error(), tt.err)
				}
			} else {
				if got != tt.expected {
					t.Errorf("got: %v, expected: %v", got, tt.expected)
				}
			}
		})
	}

}
