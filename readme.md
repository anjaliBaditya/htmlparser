# Link Parser

This package provides functionality to parse links (`<a href="...">`) from an HTML document. It extracts the href attribute and the link text.

## Installation

To install this package, use `go get`:


# Usage
Here's a simple example to demonstrate how to use the link package:

```Go
package main

import (
    "fmt"
    "strings"

    "github.com/yourusername/link"
)

func main() {
    html := `
    <html>
    <body>
        <a href="https://example.com">Example</a>
        <a href="/local">Local Link</a>
    </body>
    </html>`

    r := strings.NewReader(html)
    links, err := link.Parse(r)
    if err != nil {
        panic(err)
    }

    for _, l := range links {
        fmt.Printf("Href: %s, Text: %s\n", l.Href, l.Text)
    }
}
```

# Functions
## Parse
## #func Parse(r io.Reader) ([]Link, error)

Parse takes an HTML document as an io.Reader and returns a slice of Link structs parsed from it.