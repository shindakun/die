# die

<p align="center">
  <img style="float: right;" src="assets/die.png" alt="die gopher"/>
</p>

A simple package to roll a die or set of die for you.

die was created as a learning project.

## Installation

Use `go get github.com/shindakun/die` to install into Go.

## Usage

Import `github.com/shindakun/die`, see full example below. Now, before calling `die.Roll()` call `rand.Seed(time.Now().UTC().UnixNano())`. This is typically only necessary to call once near the start of your program. Afterward, call `die.Roll()` and pass in a string such as `"1d20"` to roll a single 20 sided die. As it stands right now you can call up to 9 die to roll at once for instance, `"9d20"` may result in `113 <nil>`.

## Usage example

```go
package main

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/shindakun/die"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  for {
    r, err := die.Roll("1d20")
    if err != nil {
      panic(err)
    }
    fmt.Println(r, err)
  }
}
```
