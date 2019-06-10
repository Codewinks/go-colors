package main

import (
	"fmt"
	"github.com/codewinks/go-colors"
)

func main() {
	colors.ShowExample()

	fmt.Println(colors.Green("Prints text in green."))
	fmt.Printf("Prints %s in blue.\n", colors.Blue("text"))

	fmt.Printf("Different colors such as %s, %s, %s, and %s.\n", colors.Blue("blue"), colors.Green("green"), colors.Red("red"), colors.Yellow("yellow"))
}
