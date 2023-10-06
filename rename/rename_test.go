package rename_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Defacto2/sceners/rename"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExampleConnect() {
	titleize := cases.Title(language.English, cases.NoLower)
	const txt = "apple and oranges"
	s := strings.Split(titleize.String(txt), " ")
	for i, w := range s {
		x := rename.Connect(w, i, len(s))
		if x != "" {
			s[i] = x
		}
	}
	fmt.Println(strings.Join(s, " "))
	// Output: Apple and Oranges
}

func ExampleFix() {
	titleize := cases.Title(language.English, cases.NoLower)
	const txt = "members of 2000ad will meet at 3pm"
	s := strings.Split(titleize.String(txt), " ")
	for i, w := range s {
		x := rename.Fix(w, i, len(s))
		if x != "" {
			s[i] = x
		}
	}
	fmt.Println(strings.Join(s, " "))
	// Output: Members of 2000AD Will Meet at 3PM
}

func ExampleFixHyphen() {
	const txt = "members-of-2000ad-will-meet-at-3pm"
	fmt.Println(rename.FixHyphen(txt))
	// Output: Members-of-2000AD-Will-Meet-at-3PM
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
		{"dot", args{"hello."}, "hello"},
		{"dots", args{"hello.."}, "hello."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.TrimDot(tt.args.s); got != tt.want {
				t.Errorf("TrimDot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmp(t *testing.T) {
	tests := []struct {
		name string
		w    string
		want string
	}{
		{"empty", "", ""},
		{"str", "hello world", "hello world"},
		{"gap amp", "hello & world", "hello & world"},
		{"gapless", "hello&world", "hello & world"},
		{"dupes", "hello&&world", "hello & world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.Amp(tt.w); got != tt.want {
				t.Errorf("Amp() = %v, want %v", got, tt.want)
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
		{"no dots", "hello.", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.Format(tt.s); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		position int
		last     int
		want     string
	}{
		{"empty", "", 0, 0, ""},
		{"first", "hello", 0, 5, ""},
		{"last", "world", 4, 5, ""},
		{"lowercase", "of", 2, 5, "of"},
		{"uppercase", "THE", 2, 5, "the"},
		{"mixedcase", "ThE", 2, 5, "the"},
		{"not a stop word", "foo", 2, 5, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rename.Connect(tt.word, tt.position, tt.last); got != tt.want {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
