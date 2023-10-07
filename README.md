# releaser

[![Go Reference](https://pkg.go.dev/badge/github.com/Defacto2/releaser.svg)](https://pkg.go.dev/github.com/Defacto2/releaser)

A Go library for handling the formatting of [Defacto2 releasers](https://defacto2.net).

Releasers are the groups or organisations that create the art, music, demos, intros, cracks, etc. that are found on the Defacto2 website.
They are also the sites and boards that have hosted the files and communities.

The library is used by the [Defacto2 website](https://defacto2.net) to format the named groups and organizations.

There are four main functions:

* `Clean` - Cleans the named releaser to correct any issues with syntax or casing.
* `Humanize` - Formats the URL path of a releaser into a human readable string.
* `Link` - Formats the URL path of a releaser for use as a HTML link description.
* `Obfuscate` - Formats the named releaser into a partial URL path.

The [name.Special](https://pkg.go.dev/github.com/Defacto2/releaser/name#Special) func contains the list of releaser names with special syntax.

The [initialism.Initialism](https://pkg.go.dev/github.com/Defacto2/releaser/initialism#Initialism) func handles the releaser alternative spellings, acronyms and initialisms.


## Usage

In your Go project, import the releaser library.

```sh
go get github.com/Defacto2/releaser
```

Use the functions.

```go
import "github.com/Defacto2/releaser"

func main() {
    // Clean the the string releaser name.
    name := releaser.Clean("  the  knightmare  bbs ")
    fmt.Println(name) // Output: Knightmare BBS

    // Format the releaser name into a URL path.
    urlPath = releaser.Obfuscate("the knightmare bbs")
    fmt.Println(urlPath) // Output: knightmare-bbs

    // Format the releaser name into a human readable string.
    const url = "https://defacto2.net/organizations/knightmare-bbs"
    name = releaser.Humanize(path.Base(url))
    fmt.Println(name) // Output: Knightmare BBS

    // Format the releaser names into a HTML link description.
    const url = "https://defacto2.net/organizations/class*paradigm*razor-1911"
    name = releaser.Link(path.Base(url))
    fmt.Println(name) // Output: Class + Paradigm + Razor 1911
}
```