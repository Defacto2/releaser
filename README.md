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
	// Clean the the string releaser name.
	name := releaser.Clean("  the  knightmare  bbs ")
	fmt.Println(name) // Output: Knightmare BBS

	// Format the releaser name for use in a database cell.
	data := releaser.Cell("  the  knightmare  bbs ")
	fmt.Println(data) // Output: KNIGHTMARE BBS

	// Format the releaser name into a URL path.
	urlPath := releaser.Obfuscate("the knightmare bbs")
	fmt.Println(urlPath) // Output: knightmare-bbs

	// Format the releaser name into a human readable string.
	const url1 = "https://defacto2.net/g/knightmare-bbs"
	name = releaser.Humanize(path.Base(url1))
	fmt.Println(name) // Output: Knightmare BBS

	// Format the releaser names into a HTML link description.
	const url2 = "https://defacto2.net/g/class*paradigm*razor-1911"
	name = releaser.Link(path.Base(url2))
	fmt.Println(name) // Output: Class + Paradigm + Razor 1911
}
```