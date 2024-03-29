package name_test

import (
	"fmt"
	"testing"

	"github.com/Defacto2/releaser/name"
)

func ExampleHumanize() {
	s, _ := name.Humanize("defacto2")
	fmt.Println(s)

	s, _ = name.Humanize("razor-1911-demo")
	fmt.Println(s)

	s, _ = name.Humanize("razor-1911-demo*trsi")
	fmt.Println(s)

	x, err := name.Humanize("razor-1911-demo#trsi")
	fmt.Printf("%q, %s\n", x, err) // print the quoted x string and any error

	// Output:
	// defacto2
	// razor 1911 demo
	// razor 1911 demo, trsi
	// "", the path contains invalid characters
}

func ExampleSpecial() {
	find := name.Path("surprise-productions")
	for key, val := range name.Special() {
		if key == find {
			fmt.Println(val)
		}
	}
	// Output: Surprise! Productions
}

func TestSpecial(t *testing.T) {
	// confirm all keys are valid and values are not empty
	special := name.Special()
	for key, val := range special {
		fmt.Println(key, val)
		if !key.Valid() {
			t.Errorf("Special() invalid %v", key)
		}
		if val == "" {
			t.Errorf("Special() empty value %v", key)
		}
	}
}

func TestHumanize(t *testing.T) {
	tests := []struct {
		name    string
		path    name.Path
		want    string
		wantErr error
	}{
		{
			name:    "valid path",
			path:    "path/to/file",
			want:    "",
			wantErr: name.ErrInvalidPath,
		},
		{
			name:    "invalid path",
			path:    "",
			want:    "",
			wantErr: name.ErrInvalidPath,
		},
		{
			name:    "path with ampersand",
			path:    "path-ampersand-path",
			want:    "path & path",
			wantErr: nil,
		},
		{
			name:    "path with underscore",
			path:    "path_with_underscore",
			want:    "path-with-underscore",
			wantErr: nil,
		},
		{
			name:    "path with asterisk",
			path:    "path*with*asterisk",
			want:    "path, with, asterisk",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := name.Humanize(tt.path)
			if err != tt.wantErr {
				t.Errorf("Humanize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Humanize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObfuscate(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "single word",
			arg:  "HeLlo",
			want: "hello",
		},
		{
			name: "multiple words",
			arg:  "Hello World",
			want: "hello-world",
		},
		{
			name: "ampersand",
			arg:  "Ben & Jerry's",
			want: "ben-ampersand-jerrys",
		},
		{
			name: "comma",
			arg:  "John, Paul, George, Ringo",
			want: "john*paul*george*ringo",
		},
		{
			name: "mixed",
			arg:  "The quick brown fox jumps over the lazy dog, but the dog is faster",
			want: "the-quick-brown-fox-jumps-over-the-lazy-dog*but-the-dog-is-faster",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := name.Obfuscate(tt.arg)
			if got != name.Path(tt.want) {
				t.Errorf("Obfuscate(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}
