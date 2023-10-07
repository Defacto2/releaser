package initialism_test

import (
	"fmt"
	"testing"

	"github.com/Defacto2/releaser/initialism"
)

func ExampleJoin() {
	s := initialism.Join("the-firm")
	fmt.Println(s) // FiRM, FRM

	s = initialism.Join("united-software-association")
	fmt.Println(s) // USA
	// Output: FiRM, FRM
	// USA
}

func TestInitialism(t *testing.T) {
	tests := []struct {
		name string
		path initialism.Path
		want []string
	}{
		{"empty path", "", nil},
		{"unknown path", "some-random-bbs", nil},
		{"known", "union", []string{"UNi"}},
		{"multiple", "wave", []string{"The Wave", "CNC"}},
		{"df2", "defacto2", []string{"DF2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialism.Initialism(tt.path); !equal(got, tt.want) {
				t.Errorf("Initialism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitialisms(t *testing.T) {
	l := initialism.Initialisms()
	if len(l) == 0 {
		t.Errorf("Initialisms() = %v, want %v", l, "non-empty")
	}
	if len(l) < 100 {
		t.Errorf("Initialisms() = %v, want %v", l, "more than 100")
	}
}

func TestIsInitialism(t *testing.T) {
	tests := []struct {
		name string
		path initialism.Path
		want bool
	}{
		{"empty path", "", false},
		{"unknown", "some-random-bbs", false},
		{"known", "tristar", true},
		{"multiple", "tristar-ampersand-red-sector-inc", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialism.IsInitialism(tt.path); got != tt.want {
				t.Errorf("IsInitialism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
