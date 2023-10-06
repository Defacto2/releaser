package sceners_test

import (
	"testing"

	"github.com/Defacto2/sceners"
)

func TestCleaner(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "The  Defacto2  Demo  Group",
			expected: "Defacto2 Demo Group",
		},
		{
			input:    "  Razor  1911  &  TRSi  ",
			expected: "Razor 1911 & TRSi",
		},
		{
			input:    "  ACiD  Productions  ",
			expected: "ACiD Productions",
		},
	}

	for _, tc := range testCases {
		actual := sceners.Cleaner(tc.input)
		if actual != tc.expected {
			t.Errorf("Cleaner(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}

func TestClean(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "/razor-1911//",
			expected: "Razor 1911",
		},
		{
			input:    "/razor-1911-ampersand-trsi",
			expected: "Razor 1911 & Trsi",
		},
		{
			input:    "/north-american-pirate_phreak-association",
			expected: "North American Pirate-Phreak Association",
		},
	}

	for _, tc := range testCases {
		actual := sceners.Clean(tc.input)
		if actual != tc.expected {
			t.Errorf("Clean(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
