# Colors ![GitHub](https://img.shields.io/github/license/codewinks/go-colors.svg) [![GoDoc](https://godoc.org/github.com/codewinks/go-colors?status.svg)](https://godoc.org/github.com/codewinks/go-colors) [![Build Status](https://img.shields.io/travis/codewinks/go-colors.svg?style=flat-square)](https://travis-ci.org/codewinks/go-colors)

Colors lets you use colorized outputs in terms of [ANSI Escape
Codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors) in Go (Golang).

Support for windows untested/coming some day.


## Install

```bash
go get github.com/codewinks/go-colors
```

## Examples
```go
import (
	"github.com/codewinks/go-colors"
)

func main(){
  colors.ShowExample()
}
```
### Standard colors

```go
	fmt.Println(colors.Green("Prints text in green."))
	fmt.Printf("Prints %s in blue.\n", colors.Blue("text"))

	fmt.Printf("Different colors such as %s, %s, %s, and %s.\n", colors.Blue("blue"), colors.Green("green"), colors.Red("red"), colors.Yellow("yellow"))
```


## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/fatih/color/blob/master/LICENSE.md) for more details
