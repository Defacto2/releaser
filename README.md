# Defacto2 / releaser

[![Go Reference](https://pkg.go.dev/badge/github.com/Defacto2/releaser.svg)](https://pkg.go.dev/github.com/Defacto2/releaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/Defacto2/magicnumber)](https://goreportcard.com/report/github.com/Defacto2/magicnumber)

A Go library for handling the formatting of Defacto2 [releasers](https://defacto2.net/releaser). See the [reference documentation](https://pkg.go.dev/github.com/Defacto2/releaser) for additional usage and examples.

Releasers are the groups or organisations that create the art, music, demos, intros, cracks, etc. that are found on the Defacto2 website. They are also the sites and boards that have hosted the files and communities.

## Usage

In your Go project, import the releaser library.

```go
package main

import (
	"fmt"
	"path"

	"github.com/Defacto2/releaser"
)

func main() {
	s := "  the  knightmare  bbs "

	// Clean the s releaser name
	// Output: Knightmare BBS
	fmt.Println(releaser.Clean(s))

	// Format the s releaser name for use in a database cell
	// Output: KNIGHTMARE BBS
	fmt.Println(releaser.Cell(s)) 


	s = "the knightmare bbs"

	// Format the s releaser name into a URL path.
	// Output: knightmare-bbs
	fmt.Println(releaser.Obfuscate(s))

	// Output: https://defacto2.net/g/knightmare-bbs
	result, _ := url.JoinPath("https://defacto2.net", "g", releaser.Obfuscate(s))
	fmt.Println(result)

	// Format the URL into a human readable string
	// Output: Knightmare BBS
	fmt.Println(releaser.Humanize(path.Base(result)))


	coop := "class*paradigm*razor-1911"

	// Format the cooperation releaser names into a HTML link description
	// Output: Class + Paradigm + Razor 1911
	result, _ := url.JoinPath("https://defacto2.net", "g", coop)
	fmt.Println(releaser.Link(path.Base(result))) 
}
```
