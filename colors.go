package colors

import (
	"fmt"
)

//NC – Reset Color
const NC = "\033[0m"

//BOLD – BOLD
const BOLD = "\033[1m"

//BLACK – BLACK
const BLACK = "\033[0;30m"

//GREY – GREY
const GREY = "\033[1;30m"

//RED – RED
const RED = "\033[0;31m"

//REDBOLD – REDBOLD
const REDBOLD = "\033[1;31m"

//GREEN – GREEN
const GREEN = "\033[0;32m"

//GREENBOLD – GREENBOLD
const GREENBOLD = "\033[1;32m"

//YELLOW – YELLOW
const YELLOW = "\033[0;33m"

//YELLOWBOLD – YELLOWBOLD
const YELLOWBOLD = "\033[1;33m"

//BLUE – BLUE
const BLUE = "\033[0;34m"

//BLUEBOLD – BLUEBOLD
const BLUEBOLD = "\033[1;34m"

//PURPLE – PURPLE
const PURPLE = "\033[0;35m"

//PURPLEBOLD – PURPLEBOLD
const PURPLEBOLD = "\033[1;35m"

//CYAN – CYAN
const CYAN = "\033[0;36m"

//CYANBOLD – CYANBOLD
const CYANBOLD = "\033[1;36m"

//WHITE – WHITE
const WHITE = "\033[0;37m"

//WHITEBOLD – WHITEBOLD
const WHITEBOLD = "\033[1;37m"

// ShowExample - Print out all different colors that are supported
func ShowExample() {
	fmt.Printf("%sNC\n", NC)
	fmt.Println(Bold("Bold"))
	fmt.Println(Black("Black"))
	fmt.Println(Grey("Grey"))
	fmt.Println(Red("Red"))
	fmt.Println(RedBold("Red Bold"))
	fmt.Println(Green("Green"))
	fmt.Println(GreenBold("Green Bold"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(YellowBold("Yellow Bold"))
	fmt.Println(Blue("Blue"))
	fmt.Println(BlueBold("Blue Bold"))
	fmt.Println(Purple("Purple"))
	fmt.Println(PurpleBold("Purple Bold"))
	fmt.Println(Cyan("Cyan"))
	fmt.Println(CyanBold("Cyan Bold"))
	fmt.Println(White("White"))
	fmt.Println(WhiteBold("White Bold"))
}

// Get – Returns formatted text with color
func Get(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, NC)
}

// Bold – Output text in Bold
func Bold(text string) string {
	return Get(BOLD, text)
}

// Black – Output text in Black
func Black(text string) string {
	return Get(BLACK, text)
}

// Grey – Output text in Grey
func Grey(text string) string {
	return Get(GREY, text)
}

// Red – Output text in Red
func Red(text string) string {
	return Get(RED, text)
}

// RedBold – Output text in RedBold
func RedBold(text string) string {
	return Get(REDBOLD, text)
}

// Green – Output text in Green
func Green(text string) string {
	return Get(GREEN, text)
}

// GreenBold – Output text in GreenBold
func GreenBold(text string) string {
	return Get(GREENBOLD, text)
}

// Yellow – Output text in Yellow
func Yellow(text string) string {
	return Get(YELLOW, text)
}

// YellowBold – Output text in YellowBold
func YellowBold(text string) string {
	return Get(YELLOWBOLD, text)
}

// Blue – Output text in Blue
func Blue(text string) string {
	return Get(BLUE, text)
}

// BlueBold – Output text in BlueBold
func BlueBold(text string) string {
	return Get(BLUEBOLD, text)
}

// Purple – Output text in Purple
func Purple(text string) string {
	return Get(PURPLE, text)
}

// PurpleBold – Output text in PurpleBold
func PurpleBold(text string) string {
	return Get(PURPLEBOLD, text)
}

// Cyan – Output text in Cyan
func Cyan(text string) string {
	return Get(CYAN, text)
}

// CyanBold – Output text in CyanBold
func CyanBold(text string) string {
	return Get(CYANBOLD, text)
}

// White – Output text in White
func White(text string) string {
	return Get(WHITE, text)
}

// WhiteBold – Output text in WhiteBold
func WhiteBold(text string) string {
	return Get(WHITEBOLD, text)
}
