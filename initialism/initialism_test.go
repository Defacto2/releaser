package initialism_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/Defacto2/releaser/initialism"
	"github.com/stretchr/testify/assert"
)

func ExampleInitialism() {
	fmt.Println(initialism.Initialism("defacto2"))
	// Output: [DF2 DF]
}

func ExampleInitialisms() {
	const find = "USA"
	for k, v := range initialism.Initialisms() {
		for _, x := range v {
			if x == find {
				fmt.Printf("Found %v in %v\n", find, k)
				return
			}
		}
	}
	// Output: Found USA in united-software-association
}

func ExampleIsInitialism() {
	fmt.Println(initialism.IsInitialism("defacto2"))
	// Output: true
}

func ExampleJoin() {
	fmt.Println(initialism.Join("the-firm")) // FiRM, FRM

	fmt.Println(initialism.Join("united-software-association")) // USA
	// Output: FiRM, FRM
	// USA
}

func TestMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"empty", "", []string{}},
		{"no match", "some-unknown-random-bbs", []string{}},
		{"df2", "df2", []string{"defacto2", "defacto2net"}},
		{"razor", "RzR", []string{"razor-1911", "razor-1911-demo", "razordox"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := initialism.Match(tt.s)
			c := make([]string, len(got))
			for i, v := range got {
				c[i] = string(v)
			}
			sort.Strings(c)
			if !assert.Equal(t, tt.want, c) {
				t.Errorf("Match() = %v, want %v", c, tt.want)
			}
		})
	}
}

func BenchmarkIsInitialism(b *testing.B) {
	b.Run("IsInitialism", func(b *testing.B) {
		fmt.Println(initialism.IsInitialism("defacto2"))
	})
}

func BenchmarkInitialism(b *testing.B) {
	b.Run("Initialism", func(b *testing.B) {
		fmt.Println(initialism.Initialism("defacto2"))
	})
}

func BenchmarkInitialisms(b *testing.B) {
	b.Run("Initialisms", func(b *testing.B) {
		const find = "USA"
		for k, v := range initialism.Initialisms() {
			for _, x := range v {
				if x == find {
					fmt.Printf("Found %v in %v\n", find, k)
					return
				}
			}
		}
	})
}

func BenchmarkMatch(b *testing.B) {
	b.Run("Match", func(b *testing.B) {
		fmt.Println(initialism.Match("razor"))
	})
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
		{"df2", "defacto2", []string{"DF2", "DF"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialism.Initialism(tt.path); !equal(got, tt.want) {
				t.Errorf("Initialism() = %v, want %v", got, tt.want)
			}
		})
	}
	// Confirm all keys are valid URL paths.
	for key := range initialism.Initialisms() {
		// keys must be lowercase and start with only letters or numbers
		k := string(key)
		chr := rune(k[0])
		assert.Equal(t, strings.ToLower(k), k)
		assert.Equal(t, strings.TrimSpace(k), k)
		assert.True(t, unicode.IsLetter(chr) || unicode.IsNumber(chr),
			"this key does not look right: "+k)
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

	s := "inc"
	m := ""
	for _, v := range l {
		for _, x := range v {
			if strings.ToLower(x) == s {
				m = x
			}
		}
	}
	if m == "" {
		t.Errorf("Initialisms() could not find %v", s)
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
