# sceners

[![Go Reference](https://pkg.go.dev/badge/github.com/Defacto2/sceners.svg)](https://pkg.go.dev/github.com/Defacto2/sceners)

A Go library for handling the formatting of [Defacto2 sceners](https://defacto2.net).

Sceners are the people, groups or organisations that create the art, music, demos, intros, cracks, etc. that are found on the Defacto2 website. They are also the sites and boards that have hosted the files and communities.

The library is used by the [Defacto2 website](https://defacto2.net) to format the sceners' names and groups.

There are two main functions:

* `Clean` - Cleans the named scener to correct syntax and casing.
* `Humanize` - Formats the URL path of a scener into a human readable string.

## Usage

```go
import "github.com/Defacto2/sceners"

func main() {
    // Clean the the string scener name.
    name := sceners.Clean("  the  knightmare  bbs ")
    fmt.Println(name) // Output: Knightmare BBS

    // Format the scener name into a human readable string.
    const url = "https://defacto2.net/organizations/knightmare-bbs"
    name = sceners.Humanize(path.Base(url)
    fmt.Println(name) // Output: Knightmare BBS
}
```