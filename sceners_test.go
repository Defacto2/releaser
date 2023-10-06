package sceners_test

import (
	"testing"

	"github.com/Defacto2/sceners"
)

func TestClean(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{""}, ""},
		{"leading the", args{"the blah"}, "The Blah"},
		{"common the", args{"in the blah"}, "In the Blah"},
		{"no spaces", args{"TheBlah"}, "Theblah"},
		{"elite fmt", args{"MiRROR now"}, "Mirror Now"},
		{"roman numbers", args{"In the row now ii"}, "In the Row Now II"},
		{"BBS", args{"MiRROR now bbS"}, "Mirror Now BBS"},
		{"slug", args{"this-is-a-slug-string"}, "This-is-a-Slug-String"},
		{
			"pair of groups",
			args{"Group inc.,RAZOR TO 1911"},
			"Group Inc,Razor to 1911",
		},
		{
			"2nd group with a leading the",
			args{"this is the group,the group is this"},
			"This is the Group,The Group is This",
		},
		{"ordinal", args{"4TH dimension"}, "4th Dimension"},
		{"ordinals", args{"4TH dimension, 5Th Dynasty"}, "4th Dimension, 5th Dynasty"},
		{"abbreviation", args{"2000 ad"}, "2000AD"},
		{"abbreviations", args{"2000ad, 500bc"}, "2000AD, 500BC"},
		{
			"mega-group",
			args{"Lightforce,Pact,TRSi,Venom,Razor 1911,the System"},
			"Lightforce,Pact,Trsi,Venom,Razor 1911,The System",
		},
		{"example help", args{"the  Defacto2  demo  group"}, "The Defacto2 Demo Group"},
		{"example help the", args{"  the x bbs  "}, "X BBS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sceners.Clean(tt.args.s); got != tt.want {
				t.Errorf("Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanize(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "defacto2",
			expected: "Defacto2",
		},
		{
			input:    "/razor-1911//",
			expected: "Razor 1911",
		},
		{
			input:    "/razor-1911-ampersand-skillion",
			expected: "Razor 1911 & Skillion",
		},
		{
			input:    "/razor-1911*trsi",
			expected: "Razor 1911, TRSi",
		},
		{
			input:    "/north-american-pirate_phreak-association",
			expected: "North American Pirate-Phreak Association",
		},
		{"2-minutes-to-midnight-bbs", "2 Minutes to Midnight BBS"},
		{"2000ad", "2000AD"},
		{"2tally-unrubbed", "2Tally Unrubbed"},
		{"2nd2none-bbs", "2ND2NONE BBS"},
		{"class*paradigm*razor-1911", "Class, Paradigm, Razor 1911"},
		{"down-town-bbs*bizare-bbs", "Down Town BBS, Bizare BBS"},
	}

	for _, tc := range testCases {
		actual := sceners.Humanize(tc.input)
		if actual != tc.expected {
			t.Errorf("Humanize(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
