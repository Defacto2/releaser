package rename_test

import (
	"testing"

	"github.com/Defacto2/sceners/rename"
)

func TestCleaner(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.Cleaner(tt.args.s); got != tt.want {
				t.Errorf("Cleaner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimThe(t *testing.T) {
	type args struct {
		g string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"The X BBS"}, "X BBS"},
		{"", args{"The X FTP"}, "X FTP"},
		{"", args{"the X BBS"}, "X BBS"},
		{"", args{"THE X BBS"}, "X BBS"},
		{"", args{"The"}, "The"},
		{"", args{"Hello BBS"}, "Hello BBS"},
		{"", args{"The High & Mighty Hello BBS"}, "High & Mighty Hello BBS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.TrimThe(tt.args.g); got != tt.want {
				t.Errorf("TrimThe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimDot(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"no dots", args{"hello"}, "hello"},
		{"dots", args{"hello."}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.TrimDot(tt.args.s); got != tt.want {
				t.Errorf("TrimDot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeObfuscate(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"2-minutes-to-midnight-bbs", "2 Minutes to Midnight BBS"},
		{"2000ad", "2000AD"},
		{"2tally-unrubbed", "2Tally Unrubbed"},
		{"2nd2none-bbs", "2ND2NONE BBS"},
		{"class*paradigm*razor-1911", "Class, Paradigm, Razor 1911"},
		{"down-town-bbs*bizare-bbs", "Down Town BBS, Bizare BBS"},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			if got := rename.DeObfuscate(tt.url); got != tt.want {
				t.Errorf("DeObfuscate() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFmtSyntax(t *testing.T) {
	tests := []struct {
		name string
		w    string
		want string
	}{
		{"empty", "", ""},
		{"str", "hello world", "hello world"},
		{"gap amp", "hello & world", "hello & world"},
		{"gapless", "hello&world", "hello & world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.FmtSyntax(tt.w); got != tt.want {
				t.Errorf("FmtSyntax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"EXACT", "beer", "BEER"},
		{"exact", "SceNET", "scenet"},
		{"specifc", "cybermail", "CyberMail"},
		{"dz", "hashx", "Hash X"},
		{"UPPER", "pcb", "PCB"},
		{"lower", "7Of9", "7of9"},
		{"exact upper", "Anz ftp", "ANZ FTP"},
		{"fmt by name", "Excretion anarchy", "eXCReTION Anarchy"},
		{"am suffix", "the 12am group", "The 12AM Group"},
		{"pm suffix", "the 12pm group", "The 12PM Group"},
		{"dox", "thedox group", "TheDox Group"},
		{"fxp", "thefxp group", "TheFXP Group"},
		{"iso", "theiso group", "TheISO Group"},
		{"nfo", "thenfo group", "TheNFO Group"},
		{"pc", "pc-group", "PC-Group"},
		{"lsd", "the lsdgroup", "The LSDGroup"},
		{"inc", "inc group", "INC Group"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.Format(tt.s); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
